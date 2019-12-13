package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		c1 <- 300
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- 100
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 200
		wg.Done()
	}()

	go func() {
		var t int
		for {
			select {
			case n := <-c1:
				t += n
			case c2 <- t:
			}
		}
	}()

	wg.Wait()
	fmt.Println(<-c2)
}
