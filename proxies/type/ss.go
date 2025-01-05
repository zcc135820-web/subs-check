package proxies

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

// SsToClash 将ss格式的节点转换为clash格式
func SsToClash(data string) (map[string]any, error) {
	if !strings.HasPrefix(data, "ss://") {
		return nil, fmt.Errorf("不是ss格式")
	}

	// 移除 "ss://" 前缀
	data = data[5:]

	// 分离名称部分
	name := ""
	if idx := strings.LastIndex(data, "#"); idx != -1 {
		name = data[idx+1:]
		name, _ = url.QueryUnescape(name)
		data = data[:idx]
	}

	// 分离用户信息和服务器信息
	parts := strings.SplitN(data, "@", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("格式错误: 缺少@分隔符")
	}

	// base64解码用户信息部分
	decoded, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("base64解码失败: %v", err)
	}

	// 分离加密方式和密码
	methodAndPassword := strings.SplitN(string(decoded), ":", 2)
	if len(methodAndPassword) != 2 {
		return nil, fmt.Errorf("格式错误: 加密方式和密码格式不正确")
	}

	method := methodAndPassword[0]
	password := methodAndPassword[1]

	// 分离服务器地址和端口
	hostPort := strings.Split(parts[1], ":")
	if len(hostPort) != 2 {
		return nil, fmt.Errorf("格式错误: 服务器地址格式不正确")
	}

	// 构建clash格式配置
	proxy := map[string]any{
		"name":     name,
		"type":     "ss",
		"server":   hostPort[0],
		"port":     hostPort[1],
		"cipher":   method,
		"password": password,
	}

	return proxy, nil
}
