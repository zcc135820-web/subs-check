package utils

import (
	"os"
	"path/filepath"

	"github.com/metacubex/mihomo/log"
)

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Errorln("获取程序路径失败: %v", err)
		return "."
	}
	return filepath.Dir(ex)
}
