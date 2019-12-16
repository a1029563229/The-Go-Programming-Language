/**
- 4.5：编写一个就地处理函数，用于去除 []string slice 中相邻的重复字符串元素
*/

package main

import "fmt"

func main() {
	a := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w", "a", "a"}
	removeMultiple(&a)
	fmt.Println(a)
}

func removeMultiple(a *[]string) {
	A := *a
	l := len(A)
	for i := 0; i < l-1; i++ {
		prev := A[i]
		next := A[i+1]
		if prev == next {
			A = append(A[:i], A[i+1:]...)
			l--
		}
	}
	*a = A
}
