package main

import(
	"fmt"
	"time"
)

func printNumbers(num int){
	for i:=1; i <= num; i++ {
		fmt.Println(i)
		//time.Sleep(time.Second)
	}
}

func printLetters(char rune){
	for i:='A'; i <= char; i++ {
		fmt.Printf("%c\n", i)
		//time.Sleep(time.Second)
	}
}

func main(){
	now := time.Now()
	defer func(){
		fmt.Println(time.Since(now))
	}()
	go printNumbers(5)
	printLetters('E')

	/* var fname string = "Ahmed"
	lname := "Mahmoud"
	for i:=1; i<4 ; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}*/

	// fmt.Println("Hello, " + fname + " " + lname)
}
