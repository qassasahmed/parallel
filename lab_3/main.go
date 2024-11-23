package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getMoney(ch chan int) {
	//Accept number of trials
	var trials int
	fmt.Println("Number of trials: ")
	if _, err := fmt.Scan(&trials); err != nil {
		fmt.Println("Please Provide Intege number")
		return
	}

	rand.Seed(time.Now().UnixNano()) //Pseudu Random
	for i := 1; i <= trials; i++ {
		amount := rand.Intn(500)
		ch <- amount
	} // End sending
	close(ch)
}

func main() {
	channel := make(chan int)
	go getMoney(channel)

	for msg := range channel { // check whether channel (sending)
		fmt.Printf("You've won: %d$\n", msg)
	}

	//	channel<-500 // Sending
	//	channel<-600
	//	fmt.Println(<-channel) // Recieving
	//	fmt.Println(<-channel)

}
