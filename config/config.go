package config

type IPInfo struct {
	APIURL  []string `yaml:"api-url"`
	IPDBURL string   `yaml:"ipdb-url"`
}

type Config struct {
	PrintProgress bool     `yaml:"print-progress"`
	Concurrent    int      `yaml:"concurrent"`
	CheckInterval int      `yaml:"check-interval"`
	DownloadSize  int      `yaml:"download-size"`
	UploadSize    int      `yaml:"upload-size"`
	Timeout       int      `yaml:"timeout"`
	FilterRegex   string   `yaml:"filter-regex"`
	SaveMethod    string   `yaml:"save-method"`
	WorkerURL     string   `yaml:"worker-url"`
	WorkerToken   string   `yaml:"worker-token"`
	SubUrls       []string `yaml:"sub-urls"`
	IPInfo        IPInfo   `yaml:"ip-info"`
}

var GlobalConfig Config
