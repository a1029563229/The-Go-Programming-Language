/**
- 1.1：修改 echo 程序输出 os.Args[0]，即命令的名称
*/

package main

import (
	"fmt"
	"os"
)

// go run main.go
func main() {
	fmt.Println(os.Args[0])
}
