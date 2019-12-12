/**
- 2.1：添加类型、常量和函数到 tempconv 包中，处理以开尔文为单位（K）的温度值， 0K = -273.15℃，变化 1K 和变化 1 ℃ 是等价的。
*/

package main

import (
	"fmt"

	"The-Go-Programming-Language/practice/2/2.1/tempconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
}
