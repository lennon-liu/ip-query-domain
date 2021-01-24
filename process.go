package dns_search

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func Process() {
	//var wp1 sync.WaitGroup
	fmt.Println(time.Now())
	var wp2 sync.WaitGroup
	var wp3 sync.WaitGroup
	inchan := make(chan *Target, config.Maxip*2)
	outchan := make(chan *Result, config.Maxip*2)
	wp3.Add(1)
	wp2.Add(config.Scanners)
	for i := 0; i < config.Scanners; i++ {

		go func(i int) {
			for target := range inchan {
				result := Httprequest(target)
				outchan <- result
			}
			wp2.Done()
		}(i)
	}

	go func() {
		i := 0
		temp_result := make([]string, 0, 1000)
		for result := range outchan {
			if result != nil {
				temp := (*result).Domain
				for _, _domain := range temp {
					if strings.Count(_domain, "")-1 > 0 {
						i++
						temp_result = append(temp_result, _domain)
					}
				}
			}
			config.WriteDomain(temp_result)
			temp_result = make([]string, 0, 1000)
		}
		wp3.Done()
	}()
	input := bufio.NewReader(config.Input)
	for {
		line, err := input.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		line_ := strings.Replace(string(line), "\r\n", "", -1)
		target := TargetPrase(line_)
		if target == nil {
			continue
		}
		if config.UseProxy {
			index := rand.Intn(len(config.Proxys) - 1)
			if &(config.Proxys[index]) != nil {
				target.ProxyURL = "http://" + config.Proxys[index].Url
			}
		}
		if target != nil {
			inchan <- target
		}
	}
	//wp1.Wait()
	close(inchan)
	wp2.Wait()
	close(outchan)
	wp3.Wait()
	fmt.Println(time.Now())
	config.Input.Close()
	config.Output.Close()
	fmt.Println("query over")
}
