/**
- 4.3：重写函数 reverse，使用数组指针作为参数而不是 slice
*/

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(&a)
	fmt.Println(a)
}

func reverse(arr *[]int) {
	A := *arr
	l := len(A)
	for i := 0; i < l/2; i++ {
		A[i], A[l-i-1] = A[l-i-1], A[i]
	}
}
