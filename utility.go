package dns_search

import "os"

type Config struct {
	Maxip     int
	Timeout   int
	Tolerant  int
	Scanners  int
	Input     *os.File
	Output    *os.File
	ProxyAddr string
	ProxyKey  string
	Proxys    []Proxy
	UseProxy bool
}

type Target struct {
	Ip       string
	ProxyURL string
}

type Result struct {
	Domain []string
	Ip     string
}

type Proxy struct {
	Url string
}

var config Config

func init() {
	config.Timeout = 5
	config.Maxip = 256
	config.Scanners = 10
	config.Tolerant = 5
	config.ProxyAddr = "127.0.0.1:6379"
	config.Proxys = make([]Proxy, 0, 300)
}
