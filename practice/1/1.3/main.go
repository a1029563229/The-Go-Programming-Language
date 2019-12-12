/**
- 1.3：尝试测量可能低效的程序和使用 strings.Join 的程序在执行时间上的差异。
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func TimeConsuming(tag string) func() {
	now := time.Now().UnixNano()
	return func() {
		after := time.Now().UnixNano()
		fmt.Printf("%q time cost %d ns\n", tag, after-now)
	}
}

func printArgs1() {
	defer TimeConsuming("printArgs1")()
	fmt.Println(os.Args[1:])
}

func printArgs2() {
	defer TimeConsuming("printArgs2")()
	args := strings.Join(os.Args[1:], " ")
	fmt.Println(args)
}

// go run main.go param1 param2
func main() {
	printArgs1() // 45000 ns
	printArgs2() // 3000 ns
}
