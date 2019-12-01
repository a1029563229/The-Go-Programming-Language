package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f == 0)
	fmt.Println(c == Celsius(f))
}
