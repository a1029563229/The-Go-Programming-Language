package main

import "fmt"

func mSlice(s []int) {
	s[0] = 100
}

func main() {
	s := []int{1, 2, 3}
	mSlice(s)
	fmt.Println(s)
}
