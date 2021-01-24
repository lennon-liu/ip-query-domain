package dns_search

import (
	"bufio"
	"github.com/go-redis/redis"
	"io"
	"os"
	"regexp"
	"strings"
)

func CheckIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func (s *Config) WriteDomain(domain []string) error {
	var result = strings.Join(domain, "\n")
	if _, err := (*s).Output.Write([]byte(result + "\n")); err != nil {
		return err
	}
	return nil
}

func (c *Config) GetProxyFredis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     (*c).ProxyAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
		Network:  "tcp",
		PoolSize: 50,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	name := (*c).ProxyKey
	val2, err := client.HGetAll(name).Result()
	for value := range val2 {
		proxy := Proxy{value}
		(*c).Proxys = append((*c).Proxys, proxy)
	}
	return nil
}

func (c *Config) GetProxyFfile(proxyfile string) error {
	pfp, err := os.Open(proxyfile)
	if err != nil {
		return err
	}
	input := bufio.NewReader(pfp)
	for {
		line, err := input.ReadBytes('\n')
		if err == io.EOF {
			//fmt.Println("end of file")
			break
		} else if err != nil {
			//fmt.Println("error of file")
			return err
		}
		line_ := strings.Replace(string(line), "\n", "", -1)
		proxy := Proxy{line_}
		(*c).Proxys = append((*c).Proxys, proxy)
	}
	return nil
}

