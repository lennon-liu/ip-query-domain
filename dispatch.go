package dns_search

import (
	"flag"
	"fmt"
	"os"
)

func ParseCommandLine(flags []string) error {
	fmt.Println(len(flags))
	var Input, Output, ProxyFile string
	flag.StringVar(&Input, "i", "", "input file path or os.stdin")
	flag.StringVar(&Output, "o", "", "output file path or os.stdout")
	flag.IntVar(&config.Tolerant, "n", 5, "proxy Tolerant")
	flag.IntVar(&config.Timeout, "t", 2, "request timeout")
	flag.IntVar(&config.Scanners, "s", 10, "request Scanners")
	flag.StringVar(&config.ProxyAddr, "pa", "", "proxytool IpAdddress fmt: 127.0.0.1:6379")
	flag.StringVar(&config.ProxyKey, "pk", "", "proxytool key fmt: proxytool")
	flag.StringVar(&ProxyFile, "pf", "", "proxytool file path ")
	flag.Parse()
	if Input != "" {
		var error_read error
		config.Input, error_read = os.Open(Input)
		if error_read != nil {
			fmt.Println(error_read)
			os.Exit(-1)
		}
	} else {
		config.Input = os.Stdin
	}

	if Output != "" {
		var error_write error
		config.Output, error_write = os.Create(Output)
		if error_write != nil {
			fmt.Println(error_write)
			os.Exit(-1)
		}
	} else {
		config.Output = os.Stdout
	}
	if ProxyFile != "" {
		if err := config.GetProxyFfile(ProxyFile); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		if len(config.Proxys) < 1 {
			panic("do not find any proxy！")
		}
		config.UseProxy=true
	} else if config.ProxyAddr != "" && config.ProxyKey != "" {
		if err := config.GetProxyFredis(); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		if len(config.Proxys) < 1 {
			panic("do not find any proxy！")
		}
		config.UseProxy=true
	}else{
		config.UseProxy=false
	}
	return nil
}

func GetConf() *Config {
	return &config
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func TargetPrase(line string) *Target {
	status := CheckIp(line)
	if status != true {
		return nil
	}
	target := new(Target)
	target.Ip = line
	return target
}
