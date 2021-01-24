package main

import "os"
import (
	"dns_search"
)

func main() {
	dns_search.ParseCommandLine(os.Args)

	dns_search.GetConf()

	dns_search.Process()
}
