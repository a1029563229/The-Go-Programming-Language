/**
- 1.8：修改 fetch 程序添加一个 http:// 前缀（假如该 URL 参数缺失协议前缀）。可能会用到 strings.HasPrefix。
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// go run main.go api.jt-gmall.com
func main() {
	for _, link := range os.Args[1:] {
		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			link = "https://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", bytes)
	}
}
