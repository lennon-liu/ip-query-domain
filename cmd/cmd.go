package main

import "os"
import _ "github.com/work/dns_search"
import (
	"github.com/work/dns_search"
)

func main() {
	dns_search.ParseCommandLine(os.Args)

	dns_search.GetConf()

	dns_search.Process()
}
