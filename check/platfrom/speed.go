package platfrom

import (
	"io"
	"net/http"
	"time"

	"github.com/bestruirui/mihomo-check/config"
	"github.com/metacubex/mihomo/log"
)

func CheckSpeed(httpClient *http.Client) (int, error) {
	// 创建一个新的测速专用客户端，基于原有客户端的传输层
	speedClient := &http.Client{
		// 设置更长的超时时间用于测速
		Timeout: 10 * time.Second,
		// 保持原有的传输层配置
		Transport: httpClient.Transport,
	}

	startTime := time.Now()

	resp, err := speedClient.Get(config.GlobalConfig.SpeedTestUrl)
	if err != nil {
		log.Debugln("测速请求失败: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	// 使用固定大小的缓冲区读取数据
	buffer := make([]byte, 32*1024) // 32KB 缓冲区
	totalBytes := 0

	for {
		n, err := resp.Body.Read(buffer)
		totalBytes += n

		if err != nil {
			if err == io.EOF {
				break
			}
			// 如果是其他错误，且已经读取了一些数据，我们仍然可以计算速度
			if totalBytes > 0 {
				break
			}
			log.Debugln("读取数据时发生错误: %v", err)
			return 0, err
		}
	}

	// 计算下载时间（毫秒）
	duration := time.Since(startTime).Milliseconds()
	if duration == 0 {
		duration = 1 // 避免除以零
	}

	// 计算速度（KB/s）
	speed := int(float64(totalBytes) / 1024 * 1000 / float64(duration))

	return speed, nil
}
