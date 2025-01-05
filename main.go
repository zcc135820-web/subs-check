package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/bestruirui/mihomo-check/check"
	"github.com/bestruirui/mihomo-check/config"
	"github.com/bestruirui/mihomo-check/proxy/ipinfo"
	"github.com/bestruirui/mihomo-check/save"
	"github.com/bestruirui/mihomo-check/utils"
	"github.com/metacubex/mihomo/log"

	"gopkg.in/yaml.v3"
)

func init() {
	configPath := flag.String("f", "", "配置文件路径")
	flag.Parse()

	execPath := utils.GetExecutablePath()

	if *configPath == "" {
		*configPath = filepath.Join(execPath, "config.yaml")
	}

	yamlFile, err := os.ReadFile(*configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Infoln("配置文件不存在，创建默认配置文件")
			err = os.WriteFile(*configPath, []byte(config.DefaultConfigTemplate), 0644)
			if err != nil {
				log.Errorln("写入默认配置文件失败: %v", err)
				os.Exit(1)
			}
			log.Infoln("默认配置文件创建成功")
			log.Infoln("请编辑配置文件: %v", *configPath)
			os.Exit(1)
		} else {
			log.Errorln("读取配置文件失败: %v", err)
			os.Exit(1)
		}
	} else {
		err = yaml.Unmarshal(yamlFile, &config.GlobalConfig)
		if err != nil {
			log.Errorln("解析配置文件失败: %v", err)
			os.Exit(1)
		}
		log.Infoln("配置文件读取成功")
	}

	ipinfo.GetIPdb()

	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",http2debug=0,httpclientlog=0")
}

func main() {

	log.Infoln("进度展示 %v", config.GlobalConfig.PrintProgress)
	interval := config.GlobalConfig.CheckInterval

	for {
		checkIP()
		nextTime := time.Now().Add(time.Duration(interval) * time.Minute)
		log.Infoln("下次检查IP时间: %v", nextTime.Format("2006-01-02 15:04:05"))
		time.Sleep(time.Duration(interval) * time.Minute)
	}

}
func checkIP() {

	log.Infoln("开始检测IP")

	results, err := check.Check()

	if err != nil {
		log.Errorln("检测IP失败: %v", err)
		os.Exit(1)
	}

	log.Infoln("检测完成")

	save.SaveConfig(results)

}
