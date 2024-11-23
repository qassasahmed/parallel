package main

import(
	"fmt"
	"time"
)

func copy_simulation(num int, done chan bool){ //Changed from printNumbers
	for i:=1; i<=num; i++{
		fmt.Printf("-")
		time.Sleep(1000 * time.Millisecond)
	}
	done<-true
}


func factorial(num int)int{
	fact := 1
	for i:=2;i<=num;i++ {
		fact *= i
	}
	return fact
}

func main(){
	now := time.Now()
	defer func(){
		fmt.Println("Execution time = ", time.Since(now))
	}()

	ch_signal := make(chan bool)

	go copy_simulation(20, ch_signal)

	var input_num int
	fmt.Printf("Enter Number: ")
	fmt.Scan(&input_num)

	answer := factorial(input_num)
	fmt.Printf("%d! = %d\n", input_num, answer)
	
	<-ch_signal
}
