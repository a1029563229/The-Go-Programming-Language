package main

import (
	"fmt"
	"sync"
	"time"
)

func run(id int) {
	time.Sleep(1 * time.Second)
	fmt.Println("run: ", id)
}

var singal = make(chan bool, 20)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			singal <- true
			defer func() {
				<-singal
				wg.Done()
			}()
			run(i)
		}(i)
	}
	wg.Wait()
}
