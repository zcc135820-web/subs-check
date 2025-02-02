package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// 将vless格式的节点转换为clash的节点
func ParseVless(data string) (map[string]any, error) {

	if !strings.HasPrefix(data, "vless://") {
		return nil, fmt.Errorf("不是vless格式")
	}

	// 移除 "vless://" 前缀
	data = strings.TrimPrefix(data, "vless://")

	// 分离用户信息和服务器信息
	parts := strings.SplitN(data, "@", 2)
	if len(parts) != 2 {
		return nil, nil
	}

	uuid := parts[0]
	remaining := parts[1]

	// 分离服务器地址和参数
	hostAndParams := strings.SplitN(remaining, "?", 2)
	if len(hostAndParams) != 2 {
		return nil, nil
	}

	// 分离服务器地址和端口
	hostPort := strings.Split(hostAndParams[0], ":")
	if len(hostPort) != 2 {
		return nil, nil
	}

	host := hostPort[0]
	port, err := strconv.Atoi(hostPort[1])
	if err != nil {
		return nil, fmt.Errorf("格式错误: 端口格式不正确")
	}

	// 解析参数
	params, err := url.ParseQuery(hostAndParams[1])
	if err != nil {
		return nil, nil
	}

	// 提取节点名称
	name := ""
	if idx := strings.LastIndex(data, "#"); idx != -1 {
		name = data[idx+1:]
		name, _ = url.QueryUnescape(name)
	}

	// 构建 clash 格式的代理配置
	proxy := map[string]any{
		"name":       name,
		"type":       "vless",
		"server":     host,
		"port":       port,
		"uuid":       uuid,
		"network":    params.Get("type"),
		"tls":        params.Get("security") == "tls",
		"servername": params.Get("sni"),
	}

	// 添加 ws 特定配置
	if params.Get("type") == "ws" {
		wsOpts := map[string]any{
			"path": params.Get("path"),
			"headers": map[string]any{
				"Host": params.Get("host"),
			},
		}
		proxy["ws-opts"] = wsOpts
	}

	return proxy, nil
}
