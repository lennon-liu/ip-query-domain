package dns_search

import (
	"fmt"
	"testing"
)

func Test_http(t *testing.T) {

	target := Target{"198.54.117.216", "http://180.101.159.206:19000"}
	result := Httprequest(&target)
	fmt.Println(*result)

}

func Test_ipcheck(t *testing.T) {
	if tesult := CheckIp("172.16.39.1888"); tesult == false {
		t.Fatalf("wrong")
	}
}
