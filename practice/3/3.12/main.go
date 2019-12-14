/**
- 3.12：编写一个函数判断两个字符串是否同文异构，也就是，它们都含有相同的字符但排列顺序不同
*/

package main

import "fmt"

func main() {
	s1 := "bdfkajbdfkabdfj"
	s2 := "asdfaghajkl;qwertyuiop["
	s3 := ";lkjhgfdsa[aapoiuytrewq"
	s4 := ";lkjhgfdsa[poiu12rewqaa"
	fmt.Println(Isomerism(s1, s2))
	fmt.Println(Isomerism(s1, s3))
	fmt.Println(Isomerism(s2, s3))
	fmt.Println(Isomerism(s2, s4))
}

type StrCount map[string]int

func Isomerism(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(StrCount)
	m2 := make(StrCount)
	for i := 0; i < len(s1); i++ {
		m1[string(s1[i])]++
		m2[string(s2[i])]++
	}

	return compareMap(m1, m2)
}

func compareMap(m1 StrCount, m2 StrCount) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, _ := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	
	return true
}
