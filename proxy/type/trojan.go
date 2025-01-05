package proxytype

import (
	"fmt"
	"net/url"
	"strings"
)

// TrojanToClash 将trojan格式的节点转换为clash格式
func TrojanToClash(data string) (map[string]any, error) {
	if !strings.HasPrefix(data, "trojan://") {
		return nil, fmt.Errorf("不是trojan格式")
	}

	// 解析URL
	u, err := url.Parse(data)
	if err != nil {
		return nil, err
	}

	// 提取密码
	password := u.User.String()

	// 分离主机和端口
	hostPort := strings.Split(u.Host, ":")
	if len(hostPort) != 2 {
		return nil, nil
	}

	// 提取节点名称
	name := ""
	if fragment := u.Fragment; fragment != "" {
		name = fragment
	}

	// 解析查询参数
	params := u.Query()

	// 构建clash格式配置
	proxy := map[string]any{
		"name":     name,
		"type":     "trojan",
		"server":   hostPort[0],
		"port":     hostPort[1],
		"password": password,
		"network":  params.Get("type"),
	}

	// 添加TLS配置
	if params.Get("security") == "tls" {
		proxy["tls"] = true
		if sni := params.Get("sni"); sni != "" {
			proxy["sni"] = sni
		}
	}

	// 根据不同传输方式添加特定配置
	switch params.Get("type") {
	case "ws":
		wsOpts := map[string]any{
			"path": params.Get("path"),
		}
		if host := params.Get("host"); host != "" {
			wsOpts["headers"] = map[string]any{
				"Host": host,
			}
		}
		proxy["ws-opts"] = wsOpts
	case "grpc":
		if serviceName := params.Get("serviceName"); serviceName != "" {
			proxy["grpc-opts"] = map[string]any{
				"serviceName": serviceName,
			}
		}
	}

	return proxy, nil
}
