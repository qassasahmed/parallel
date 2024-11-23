package main

import (
	"fmt"
	"sync"
)

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup this Goroutine is done
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)              // Increment WaitGroup counter
		go printNumber(i, &wg) // Start Goroutine with WaitGroup
	}
	wg.Wait() // Wait for all Goroutines to complete
}
