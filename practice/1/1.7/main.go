/**
- 1.7：函数 io.Copy(dst, src) 从 src 读，并且写入 dst。使用它代替 ioutil.ReadAll 来复制响应内容到 os.Stdout，
这样不需要装下整个响应数据流的缓冲区。确保检查 io.Copy 返回的错误结果。
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// go run main.go https://www.baidu.com
func main() {
	for i, link := range os.Args[1:] {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		filename := fmt.Sprintf("file%d", i)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			log.Fatal(err)
		}
	}
}
