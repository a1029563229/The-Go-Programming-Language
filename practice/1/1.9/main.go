/**
- 1.9：修改 fetch 来输出 HTTP 的状态码，可以在 resp.Status 中找到它。
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// go run main.go www.baidu.com
func main() {
	for _, link := range os.Args[1:] {
		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			link = "http://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}
}
