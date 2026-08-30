package main 

import "fmt"

// function that calculated the factorial of the number 

func fact(a int, ch chan int){

	f:= 1

	for i:= 2; i <= a; i++{
		f *= i
	}
	// storing the value in the channel

	ch <- f
}

func main(){

	// make here one channel to receive the value

	r:= make(chan int)

	// n := 5

	go fact(5,r)
	
	// save the value

	ty := <-r

	fmt.Println(ty)

	// Using for loop we can print all the factorial of the form  1 to 120 

	d := make(chan int)
	for i:= 1; i <= 20 ; i++{

		go fact(i,d)

		gy:= <- d 
		fmt.Println(gy)
	}

}