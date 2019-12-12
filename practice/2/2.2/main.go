/**
- 2.2：写一个类似于 cf 的通用的单位转换程序，从命令行参数或者标准输入（如果没有参数）获取数字，
然后将每一个数字转换为以摄氏温度和华氏温度表示的温度，以英寸和米表示的长度单位，以磅和千克表示的重量，等等。
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"The-Go-Programming-Language/practice/2/2.2/tempconv"
)

func main() {
	for _, str := range os.Args[1:] {
		t, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
		}
		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)

		fmt.Printf("%s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))

		in := tempconv.Inch(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", in, tempconv.IToM(in), m, tempconv.MToI(m))

		ib := tempconv.Ib(t)
		k := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", ib, tempconv.IToK(ib), k, tempconv.KToI(k))
	}
}
