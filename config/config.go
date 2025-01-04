package config

type Config struct {
	Concurrent    int      `yaml:"concurrent"`
	CheckInterval int      `yaml:"check-interval"`
	DownloadSize  int      `yaml:"download-size"`
	UploadSize    int      `yaml:"upload-size"`
	Timeout       int      `yaml:"timeout"`
	FilterRegex   string   `yaml:"filter-regex"`
	SubUrls       []string `yaml:"sub-urls"`
	WorkerURL     string   `yaml:"worker-url"`
	WorkerToken   string   `yaml:"worker-token"`
	PrintProgress bool     `yaml:"print-progress"`
	IPInfo        struct {
		APIURL  []string `yaml:"api-url"`
		IPDBURL string   `yaml:"ipdb-url"`
	} `yaml:"ip-info"`
}

var GlobalConfig Config
