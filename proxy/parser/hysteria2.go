package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func ParseHysteria2(data string) (map[string]any, error) {
	if !strings.HasPrefix(data, "hysteria2://") && !strings.HasPrefix(data, "hy2://") {
		return nil, fmt.Errorf("不是hysteria2格式")
	}

	// 移除 "hysteria2://" 前缀

	link, err := url.Parse(data)
	if err != nil {
		return nil, err
	}

	username := link.User.Username()
	password, exist := link.User.Password()
	if !exist {
		password = username
	}
	query := link.Query()
	server := link.Hostname()
	if server == "" {
		return nil, fmt.Errorf("hysteria2 服务器地址错误")
	}
	portStr := link.Port()
	if portStr == "" {
		return nil, fmt.Errorf("hysteria2 端口错误")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("hysteria2 端口错误")
	}
	network, obfs, obfsPassword, pinSHA256, insecure, sni := query.Get("network"), query.Get("obfs"), query.Get("obfs-password"), query.Get("pinSHA256"), query.Get("insecure"), query.Get("sni")
	enableTLS := pinSHA256 != "" || sni != ""
	insecureBool := insecure == "1"

	return map[string]any{
		"type":           "hysteria2",
		"name":           username,
		"server":         server,
		"port":           port,
		"password":       password,
		"obfs":           obfs,
		"obfsParam":      obfsPassword,
		"sni":            sni,
		"skipCertVerify": insecureBool,
		"tls":            enableTLS,
		"network":        network,
	}, nil
}
