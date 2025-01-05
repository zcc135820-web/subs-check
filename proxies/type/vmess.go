package proxies

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// VmessToClash 将vmess格式的节点转换为clash格式
func VmessToClash(data string) (map[string]any, error) {
	if !strings.HasPrefix(data, "vmess://") {
		return nil, fmt.Errorf("不是vmess格式")
	}
	// 移除 "vmess://" 前缀
	data = data[8:]

	// base64解码
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// 解析JSON
	var vmessInfo map[string]string
	if err := json.Unmarshal(decoded, &vmessInfo); err != nil {
		return nil, err
	}

	// 构建clash格式配置
	proxy := map[string]any{
		"name":       vmessInfo["ps"],
		"type":       "vmess",
		"server":     vmessInfo["add"],
		"port":       vmessInfo["port"],
		"uuid":       vmessInfo["id"],
		"alterId":    vmessInfo["aid"],
		"cipher":     "auto",
		"network":    vmessInfo["net"],
		"tls":        vmessInfo["tls"] == "tls",
		"servername": vmessInfo["sni"],
	}

	// 根据不同传输方式添加特定配置
	switch vmessInfo["net"] {
	case "ws":
		wsOpts := map[string]any{
			"path": vmessInfo["path"],
		}
		if host := vmessInfo["host"]; host != "" {
			wsOpts["headers"] = map[string]any{
				"Host": host,
			}
		}
		proxy["ws-opts"] = wsOpts
	case "grpc":
		grpcOpts := map[string]any{
			"serviceName": vmessInfo["path"],
		}
		proxy["grpc-opts"] = grpcOpts
	}

	// 添加 ALPN 配置
	if alpn := vmessInfo["alpn"]; alpn != "" {
		proxy["alpn"] = strings.Split(alpn, ",")
	}

	return proxy, nil
}
