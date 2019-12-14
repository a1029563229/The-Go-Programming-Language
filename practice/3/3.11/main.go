/**
- 3.11：增强 comma 函数的功能，使其正确处理浮点数，以及带有可选正负号的数字
*/

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s1 := "832921391279123"
	s2 := "+832921391279123"
	s3 := "+8329213912.1512"
	s4 := "-82.1124123"
	fmt.Println(comma(s1))
	fmt.Println(comma(s2))
	fmt.Println(comma(s3))
	fmt.Println(comma(s4))
}

func comma(s string) string {
	var intStr string
	var floatStr string
	var prefix string

	intStr = s
	// 判断正负号
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		prefix = s[:1]
		intStr = s[1:]
	}

	// 判断浮点数
	if strings.Contains(intStr, ".") {
		combine := strings.Split(intStr, ".")
		intStr = combine[0]
		floatStr = "." + combine[1]
	}

	if len(intStr) <= 3 {
		return s
	}

	strBuffer := bytes.NewBufferString(prefix)
	strBuffer.WriteString(intStr[:3])
	for i, b := range intStr[3:] {
		if i%3 == 0 {
			strBuffer.WriteString(",")
		}
		strBuffer.WriteRune(b)
	}
	strBuffer.WriteString(floatStr)
	return strBuffer.String()
}
