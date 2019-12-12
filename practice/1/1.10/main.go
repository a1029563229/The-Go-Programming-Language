/**
- 1.10：找一个产生大量数据的网站。连续两次运行 fetchall，看报告的时间是否会有大的变化，调查缓存情况。
每一次获取的内容一样吗？修改 fetchall 将内容输出到文件，这样可以检查它是否一致。
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

// go run main.go www.jd.com
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
