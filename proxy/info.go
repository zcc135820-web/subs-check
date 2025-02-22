package proxies

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func GetProxyCountry(httpClient *http.Client) string {

	infoClient := &http.Client{
		// 设置更长的超时时间用于获取代理国家
		Timeout: time.Duration(10) * time.Second,
		// 保持原有的传输层配置
		Transport: httpClient.Transport,
	}
	// 定义多个 IP 查询 API
	apis := []string{
		"https://api.ip.sb/geoip",
		"https://ipapi.co/json",
		"https://ip.seeip.org/geoip",
		"https://api.myip.com",
	}

	for _, api := range apis {
		for attempts := 0; attempts < 5; attempts++ {
			req, err := http.NewRequest("GET", api, nil)
			if err != nil {
				time.Sleep(time.Second * time.Duration(attempts))
				continue
			}

			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
			req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Pragma", "no-cache")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")
			req.Header.Set("Sec-Ch-Ua", `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`)
			req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
			req.Header.Set("Sec-Ch-Ua-Platform", "Windows")
			req.Header.Set("Sec-Fetch-Dest", "document")
			req.Header.Set("Sec-Fetch-Mode", "navigate")
			req.Header.Set("Sec-Fetch-Site", "none")
			req.Header.Set("Sec-Fetch-User", "?1")
			req.Header.Set("Upgrade-Insecure-Requests", "1")

			resp, err := infoClient.Do(req)
			if err != nil {
				time.Sleep(time.Second * time.Duration(attempts))
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				time.Sleep(time.Second * time.Duration(attempts))
				continue
			}

			ipinfo := map[string]any{}
			err = json.Unmarshal(body, &ipinfo)
			if err != nil {
				time.Sleep(time.Second * time.Duration(attempts))
				continue
			}

			// 不同 API 返回的国家代码字段名可能不同
			countryCode := ""
			ok := false
			switch api {
			case "https://api.ip.sb/geoip":
				if code, exists := ipinfo["country_code"].(string); exists {
					countryCode = code
					ok = true
				}
			case "https://ipapi.co/json":
				if code, exists := ipinfo["country_code"].(string); exists {
					countryCode = code
					ok = true
				}
			case "https://ip.seeip.org/geoip":
				if code, exists := ipinfo["country_code"].(string); exists {
					countryCode = code
					ok = true
				}
			case "https://api.myip.com":
				if code, exists := ipinfo["cc"].(string); exists {
					countryCode = code
					ok = true
				}
			}

			if ok && countryCode != "" {
				return countryCode
			}
		}
	}
	return ""
}
