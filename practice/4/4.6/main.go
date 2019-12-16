/**
- 4.6：编写一个就地处理函数，用于将一个 UTF-8 编码的字节 slice 中所有相邻的 Unicode 空白字符
（查看 unicode.IsSpace）缩减为一个 ASCII 空白字符
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "asdf    sadf  sd f a   sfd    a"
	removeEmpty(&a)
	fmt.Println(a)
}

func removeEmpty(a *string) {
	S := *a
	str := string(S[0])
	end := 0
	for _, s := range S {
		last := rune(str[end])
		if unicode.IsSpace(last) && unicode.IsSpace(s) {
			continue
		}
		str += string(s)
		end++
	}
	*a = str[1:]
}
