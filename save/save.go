package save

import (
	"github.com/bestruirui/mihomo-check/check"
	"github.com/metacubex/mihomo/log"
	"gopkg.in/yaml.v3"
)

func SaveConfig(results []check.Result) {

	all := make([]map[string]any, 0)
	openai := make([]map[string]any, 0)
	google := make([]map[string]any, 0)
	cloudflare := make([]map[string]any, 0)
	youtube := make([]map[string]any, 0)
	netflix := make([]map[string]any, 0)
	disney := make([]map[string]any, 0)
	for _, result := range results {
		all = append(all, result.Proxy)
		if result.Openai {
			openai = append(openai, result.Proxy)
		}
		if result.Google {
			google = append(google, result.Proxy)
		}
		if result.Cloudflare {
			cloudflare = append(cloudflare, result.Proxy)
		}
		if result.Youtube {
			youtube = append(youtube, result.Proxy)
		}
		if result.Netflix {
			netflix = append(netflix, result.Proxy)
		}
		if result.Disney {
			disney = append(disney, result.Proxy)
		}
	}

	yamlData, err := yaml.Marshal(map[string]any{
		"proxies": all,
	})
	if err != nil {
		log.Errorln("解析 all 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "all")
	if err != nil {
		log.Errorln("上传 all 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": openai,
	})
	if err != nil {
		log.Errorln("解析 openai 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "openai")
	if err != nil {
		log.Errorln("上传 openai 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": google,
	})
	if err != nil {
		log.Errorln("解析 google 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "google")
	if err != nil {
		log.Errorln("上传 google 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": cloudflare,
	})
	if err != nil {
		log.Errorln("解析 cloudflare 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "cloudflare")
	if err != nil {
		log.Errorln("上传 cloudflare 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": youtube,
	})
	if err != nil {
		log.Errorln("解析 youtube 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "youtube")
	if err != nil {
		log.Errorln("上传 youtube 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": netflix,
	})
	if err != nil {
		log.Errorln("解析 netflix 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "netflix")
	if err != nil {
		log.Errorln("上传 netflix 失败: %v\n", err)
	}

	yamlData, err = yaml.Marshal(map[string]any{
		"proxies": disney,
	})
	if err != nil {
		log.Errorln("解析 disney 失败: %v\n", err)
	}
	err = UploadToKV(yamlData, "disney")
	if err != nil {
		log.Errorln("上传 disney 失败: %v\n", err)
	}

}
