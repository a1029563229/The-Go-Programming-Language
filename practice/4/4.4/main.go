/**
- 4.4：编写一个函数 rotate，实现一次遍历就可以完成元素旋转。
*/

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := rotate(a, 3)
	fmt.Println(r)
}

func rotate(s []int, position int) []int {
	r := s[position:]
	for i := position - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
