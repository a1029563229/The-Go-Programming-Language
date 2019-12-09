package main

import (
	"fmt"
	"time"
)

func runTask(taskId int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(taskId)
}

func main() {
	task := make(chan int)
	for i := 0; i < 100; i++ {
		go func(taskId int) {
			runTask(taskId)
			task <- taskId
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-task
	}
}
