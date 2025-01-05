package proxies

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 订阅链接中获取数据
func GetDateFromSubs(subUrl string) ([]byte, error) {
	maxRetries := 30
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			time.Sleep(time.Second)
		}
		resp, err := http.Get(subUrl)
		if err != nil {
			lastErr = err
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			continue
		}
		return body, nil
	}

	return nil, fmt.Errorf("重试%d次后失败: %v", maxRetries, lastErr)
}

// 从订阅链接中的数据判断是否是base64,如果是base64则解码,否则直接返回
func IsBase64(data []byte) (bool, []byte, error) {
	if strings.HasPrefix(string(data), "base64://") {
		decoded, err := base64.StdEncoding.DecodeString(string(data)[7:])
		if err != nil {
			return false, nil, err
		}
		return true, decoded, nil
	}
	return false, data, nil
}

// func GetProxies(data []byte) ([]map[string]any, error) {
// }
