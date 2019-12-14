/**
- 3.10：编写一个非递归的 comma 函数，运用 bytes.Buffer，而不是简单的字符串拼接
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "832921391279123"
	fmt.Println(comma(s))
	fmt.Println(comma2(s))
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	strBuffer := bytes.NewBufferString(s[:3])
	for i, b := range s[3:] {
		if i%3 == 0 {
			strBuffer.WriteString(",")
		}
		strBuffer.WriteRune(b)
	}
	return strBuffer.String()
}

func comma2(s string) string {
	if len(s) <= 3 {
		return s
	}

	var newStr string
	strBuffer := bytes.NewBufferString(s)
	for {
		b := strBuffer.Next(3)
		if len(b) == 0 {
			break
		}
		newStr = newStr + string(b) + ","
	}
	return newStr[:len(newStr)-1]
}
