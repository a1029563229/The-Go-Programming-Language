/**
- 2.3：使用循环重写 PopCount 来代替单个表达式。对比两个版本的效率。
*/

package main

import (
	"fmt"

	"The-Go-Programming-Language/practice/2/2.3/popcount"
)

func main() {
	fmt.Println(popcount.PopCount1(200)) // "PopCount1" time cost 2000 ns
	fmt.Println(popcount.PopCount2(200)) // "PopCount2" time cost 0 ns
}
