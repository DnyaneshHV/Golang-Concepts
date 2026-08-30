package main

import "fmt"

func main(){

	add()
	whilePractice()

}

func add() {

	a :=  6
	b :=  6

	c := a + b

	fmt.Println("Addition of two numers is : ", c)
}

func whilePractice() { 

     i := 1
	// here in go language the while is written like the syntax of while is like for

	for i<=5 {
		fmt.Println("Number : " , i)
		i++
	}
}