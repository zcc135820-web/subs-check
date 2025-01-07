package method

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bestruirui/mihomo-check/config"
	"github.com/metacubex/mihomo/log"
)

const (
	maxRetries    = 3
	retryInterval = 2 * time.Second
)

// KVPayload 定义上传到R2的数据结构
type KVPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// R2Uploader 处理R2存储上传的结构体
type R2Uploader struct {
	client    *http.Client
	workerURL string
	token     string
}

// NewR2Uploader 创建新的R2上传器
func NewR2Uploader() *R2Uploader {
	return &R2Uploader{
		client:    &http.Client{Timeout: 30 * time.Second},
		workerURL: config.GlobalConfig.WorkerURL,
		token:     config.GlobalConfig.WorkerToken,
	}
}

// UploadToR2Storage 上传数据到R2存储的入口函数
func UploadToR2Storage(yamlData []byte, key string) error {
	uploader := NewR2Uploader()
	return uploader.Upload(yamlData, key)
}

// Upload 执行上传操作
func (r *R2Uploader) Upload(yamlData []byte, key string) error {
	// 验证输入
	if err := r.validateInput(yamlData, key); err != nil {
		return err
	}

	// 准备请求数据
	payload := KVPayload{
		Key:   key,
		Value: string(yamlData),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON编码失败: %w", err)
	}

	// 执行带重试的上传
	return r.uploadWithRetry(jsonData, key)
}

// validateInput 验证输入参数
func (r *R2Uploader) validateInput(yamlData []byte, key string) error {
	if len(yamlData) == 0 {
		return fmt.Errorf("yaml数据为空")
	}
	if key == "" {
		return fmt.Errorf("key不能为空")
	}
	if r.workerURL == "" || r.token == "" {
		return fmt.Errorf("Worker配置不完整")
	}
	return nil
}

// uploadWithRetry 带重试机制的上传
func (r *R2Uploader) uploadWithRetry(jsonData []byte, key string) error {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if err := r.doUpload(jsonData); err != nil {
			lastErr = err
			log.Errorln("上传失败(尝试 %d/%d): %v", attempt+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}
		log.Infoln("上传成功: %s", key)
		return nil
	}

	return fmt.Errorf("上传失败，已重试%d次: %w", maxRetries, lastErr)
}

// doUpload 执行单次上传
func (r *R2Uploader) doUpload(jsonData []byte) error {
	// 创建请求
	req, err := r.createRequest(jsonData)
	if err != nil {
		return err
	}

	// 发送请求
	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应
	return r.checkResponse(resp)
}

// createRequest 创建HTTP请求
func (r *R2Uploader) createRequest(jsonData []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s/storage?token=%s", r.workerURL, r.token)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// checkResponse 检查响应结果
func (r *R2Uploader) checkResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取响应失败(状态码: %d): %w", resp.StatusCode, err)
		}
		return fmt.Errorf("上传失败(状态码: %d): %s", resp.StatusCode, string(body))
	}
	return nil
}
