/**
- 1.11：使用更长的参数列表来尝试 fetchall，例如使用 alexa.com 排名前 100 万的网站。如果一个网站没有响应，程序的行为是怎样的？
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano() / 1000000
	return func() {
		after := time.Now().UnixNano() / 1000000
		fmt.Printf("%q time cost %d ms\n", tag, after-now)
	}
}

// go run main.go api.jt-gmall.com
func main() {
	ch := make(chan string)
	for _, link := range os.Args[1:] {
		go fetch(link, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(link string, ch chan<- string) {
	defer TimeConsuming("fetch " + link)()
	if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
		link = "http://" + link
	}

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- string(bytes)
}
