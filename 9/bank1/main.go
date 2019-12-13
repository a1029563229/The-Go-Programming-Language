package main

import (
	"The-Go-Programming-Language/9/bank1/bank"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		wg.Done()
	}()

	go func() {
		bank.Deposit(100)
		wg.Done()
	}()

	wg.Wait()
}
