package proxies

import (
	"fmt"
	"net"
)

func DeduplicateProxies(proxies []map[string]any) []map[string]any {
	// 使用map来存储唯一的代理配置
	seen := make(map[string]map[string]any)

	// 遍历所有代理配置
	for _, proxy := range proxies {
		// 获取server和port值
		server, serverOk := proxy["server"].(string)
		port, portOk := proxy["port"].(int)
		// 如果server或port不存在，保留该配置
		if !serverOk || !portOk {
			continue
		}
		//查询server的ip
		serverip, err := net.LookupIP(server)
		if err != nil {
			continue
		}

		// 创建唯一键
		key := fmt.Sprintf("%s:%v", serverip, port)

		// 如果这个组合之前没有出现过，将其添加到seen map中
		if _, exists := seen[key]; !exists {
			seen[key] = proxy
		}
	}

	// 将去重后的配置转换回切片
	result := make([]map[string]any, 0, len(seen))
	for _, proxy := range seen {
		result = append(result, proxy)
	}

	return result
}
