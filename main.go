package main

import (
	"os"
	"time"

	"github.com/bestruirui/mihomo-check/check"
	"github.com/bestruirui/mihomo-check/config"
	"github.com/bestruirui/mihomo-check/ipinfo"
	"github.com/bestruirui/mihomo-check/save"
	"github.com/metacubex/mihomo/log"

	"gopkg.in/yaml.v3"
)

var checker *check.Check

func init() {
	version, err := os.ReadFile("/app/version")
	if err != nil {
		log.Errorln("读取版本文件失败: %v", err)
		os.Exit(1)
	}
	log.Infoln("构建时间: %s", string(version))
	yamlFile, err := os.ReadFile("/app/config/config.yaml")
	if err != nil {
		log.Errorln("读取配置文件失败: %v", err)
		os.Exit(1)
	}
	yaml.Unmarshal(yamlFile, &config.GlobalConfig)
	log.Infoln("配置文件读取成功")
	ipinfo.GetIPdb()
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",http2debug=0,httpclientlog=0")
}

func main() {

	log.Infoln("进度展示 %v", config.GlobalConfig.PrintProgress)
	log.Infoln("线程数量 %v", config.GlobalConfig.Concurrent)
	interval := config.GlobalConfig.CheckInterval
	checker = check.New()

	for {
		checkIP()
		// 打印下次检查IP时间
		nextTime := time.Now().Add(time.Duration(interval) * time.Minute)
		log.Infoln("下次检查IP时间: %v", nextTime.Format("2006-01-02 15:04:05"))
		time.Sleep(time.Duration(interval) * time.Minute)
	}

}
func checkIP() {

	log.Infoln("开始检测IP")

	err := checker.Start()

	if err != nil {
		log.Errorln("检测失败: %v", err)
		return
	}

	log.Infoln("检测完成")

	results := checker.GetResults()

	save.SaveConfig(results)

}
