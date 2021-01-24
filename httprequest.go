package dns_search

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"regexp"
	"time"
)

func Httprequest(target *Target) (result *Result) {
	result = new(Result)
	result.Ip = target.Ip
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if target.ProxyURL != "" {
		if proxy, err := u.Parse(target.ProxyURL); err != nil {
			panic(err)
		} else {
			transport.Proxy = http.ProxyURL(proxy)
		}
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(time.Duration(config.Timeout) * time.Second),
	}
	url := fmt.Sprintf("https://site.ip138.com/%s/", target.Ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		return result
	}
	if response != nil {
		defer response.Body.Close()
	}
	body, readerr := ioutil.ReadAll(response.Body)
	if readerr != nil {
		return nil
	}
	pattern := "[\\d-]+-----[\\d-]+</span><a href=\"/([a-z0-9A-Z\\-]+.+[a-z0-9A-Z])/\""
	re := regexp.MustCompile(pattern)
	matchs := re.FindAllStringSubmatch(string(body), config.Tolerant)
	for _, match := range matchs {
		result.Domain = append(result.Domain, match[1])
	}
	return result
}
