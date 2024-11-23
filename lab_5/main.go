package main

import	"fmt"

/* main(){
	number := [5]int{1, 2, 3, 4, 5}
	fmt.Println(number) // [1, 2, 3, 4, 5]

	slice := number[2:4] // len = 2, cap = 3
	fmt.Print("Cap: ", cap(slice))

	slice[0] = -10
	fmt.Println("new Slice", slice) // [-10, 4] 
	fmt.Println("new Array", number) //[1, 2, -10, 4, 5]
}*/

func main(){

	type Student struct {
		courses [5]string
		grades [5]float64
	}

	qassas := Student{
		courses:[5]string{"DB", "IT", "AI", "ML", "CN"},
		grades:[5]float64{85.0, 90.0, 92.0, 95.5, 100.0},
	}

	fmt.Println("course: ", qassas.courses)

}
