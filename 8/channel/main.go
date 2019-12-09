package main

import (
	"fmt"
	"time"
)

func run(c chan<- bool) {
	time.Sleep(1 * time.Second)
	c <- true
	fmt.Println("run")
}

func main() {
	c := make(chan bool)
	go run(c)
	select {
	case <-c:
		fmt.Println("do something")
	}

	fmt.Println("over")
}
