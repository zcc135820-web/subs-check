package save

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bestruirui/mihomo-check/config"
	"github.com/metacubex/mihomo/log"
)

type KVPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func UploadToKV(yamlData []byte, key string) error {
	// 创建请求负载
	payload := KVPayload{
		Key:   key,
		Value: string(yamlData),
	}

	// 转换为JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON编码失败: %v", err)
	}

	// 创建POST请求
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/kv?token=%s", config.GlobalConfig.WorkerURL, config.GlobalConfig.WorkerToken),
		bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取响应失败: %v", err)
		}
		return fmt.Errorf("上传失败: %s", string(body))
	}
	log.Infoln("上传成功: %s\n", key)
	return nil
}
