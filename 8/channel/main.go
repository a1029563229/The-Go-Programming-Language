package main

import "fmt"

func run(c chan<- bool) {
	close(c)
	fmt.Println("run")
}

func main() {
	c := make(chan bool)
	go run(c)
	fmt.Println(<-c)
	fmt.Println("over")
}
