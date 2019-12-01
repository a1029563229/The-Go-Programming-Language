package main

import (
	"The-Go-Programming-Language/ProgrammingStructure/currencyconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		r := currencyconv.RMB(t)
		h := currencyconv.HongKongDollar(t)
		fmt.Printf("%s = %s, %s = %s\n",
			r, currencyconv.RToH(r), h, currencyconv.HToR(h))
	}
}
