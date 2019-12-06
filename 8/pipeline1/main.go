package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			time.Sleep(500 * time.Millisecond)
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
