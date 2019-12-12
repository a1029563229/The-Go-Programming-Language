/**
- 1.2：修改 echo 程序，输出参数的索引和值，每行一个。
*/

package main

import (
	"fmt"
	"os"
)

// go run main.go param1 param2
func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Println(i)
		fmt.Println(args[i])
	}
}
