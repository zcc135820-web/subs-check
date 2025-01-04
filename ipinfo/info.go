package ipinfo

import (
	"io"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/bestruirui/mihomo-check/config"
	"github.com/metacubex/mihomo/log"
)

var (
	cityDB *City
	once   sync.Once
)

// 初始化数据库的函数
func initDB() {
	db, err := NewCity("/app/openipdb.ipdb")
	if err != nil {
		log.Errorln("初始化IP数据库失败: %v", err)
		return
	}
	cityDB = db
}
func GetIPaddrFromDNS(server string) string {
	ns, err := net.LookupIP(server)
	if err != nil {
		log.Errorln("获取IP地址失败: %v", err)
		return ""
	}
	for _, ip := range ns {
		if ip.To4() != nil {
			return ip.String()
		}
	}
	log.Errorln("server: %s, ip: %s", server, ns)
	return ""
}

func GetIPaddrFromAPI(httpClient *http.Client) string {
	for _, u := range config.GlobalConfig.IPInfo.APIURL {
		resp, err := httpClient.Get(u)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		ip := string(body)
		// 确保是ipv4
		if ip4 := net.ParseIP(ip).To4(); ip4 == nil {
			continue
		}
		return ip
	}
	return ""
}

type IPResponse struct {
	CountryName string `json:"country_name"`
}

func GetIPCountrynameFromdb(ip string) string {

	once.Do(initDB)

	if cityDB == nil {
		return ""
	}

	country, err := cityDB.Find(ip, "CN")
	if err != nil {
		return ""
	}

	return country[0]
}

func GetIPdb() {
	// 判断文件是否存在,存在则跳过
	if _, err := os.Stat("/app/openipdb.ipdb"); err == nil {
		log.Infoln("IP数据库已存在")
		return
	}
	log.Infoln("IP数据库不存在,开始下载")
	resp, err := http.Get(config.GlobalConfig.IPInfo.IPDBURL)
	if err != nil {
		log.Errorln("获取IP数据库失败: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("读取IP数据库失败: %v", err)
		return
	}
	os.WriteFile("/app/openipdb.ipdb", body, 0644)
	log.Infoln("IP数据库下载成功")
}
