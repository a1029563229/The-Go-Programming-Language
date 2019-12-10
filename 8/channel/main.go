package main

import (
	"fmt"
	"time"
)

func run(c chan<- bool) {
	time.Sleep(1 * time.Second)
	c <- true
	c <- false
	fmt.Println("run")
}

func main() {
	c := make(chan bool)
	go run(c)
	// <-c
	<-c

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("over")
}
