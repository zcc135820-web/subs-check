package proxies

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func GetProxyCountry(httpClient *http.Client) string {
	for attempts := 0; attempts < 4; attempts++ {
		req, err := http.NewRequest("GET", "https://api.ip.sb/geoip", nil)
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

		resp, err := httpClient.Do(req)
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
		return ipinfo["country_code"].(string)
	}
	return ""
}
