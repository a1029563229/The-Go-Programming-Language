/**
- 4.2：编写一个程序，用于在默认情况下输出其标准输入的 SHA256 散列，但也支持一个输出 SHA384 或 SHA512 散列的命令行标记
*/

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"os"
	"strconv"
)

// go run main.go 384 hello
func main() {
	bs := os.Args[1]
	s := os.Args[2]
	b, _ := strconv.ParseInt(bs, 10, 64)
	fmt.Println(encrypt(s, b))
}

func encrypt(s string, b int64) string {
	var h hash.Hash
	switch b {
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		h = sha256.New()
	}
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
