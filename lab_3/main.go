package main

import(
	"fmt"
	"time"
	"math/rand"
)

func getMoney(channel chan int){
//      Accept input
	var trials int
	fmt.Printf("Number of trials: ")
	if _, err := fmt.Scan(&trials); err != nil{
		fmt.Println("\nPlease Enter an Integer Value")
		return
	}

//      Iterating through the channel(send)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= trials; i++ {
		amount := rand.Intn(5000)
		channel <- amount
	}
//      Make sure to close the channel to avoid deadlocks	
	close(channel)
}

func main(){
	channel := make(chan int)
	go getMoney(channel)

//      Iterate through the channel (recieve)
	for msg := range channel{
		fmt.Printf("\nYou've won: %d EGP\n", msg)
	}
	

/*	for i := 1; i <= trials; i++ {
		go getMoney(channel)
		fmt.Printf("\nTrial %d You've won: %d EGP\n", i, <- channel)
	}
*/
}
