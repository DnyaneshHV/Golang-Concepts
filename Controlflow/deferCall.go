package main

import "fmt"

func main(){

	

	fmt.Println("Hellow World")

	defer add()  // defer make this function to wait and first execute the around functions and then it executes.

	mul()

}

func add() {

	a := 5
	b := 4
	c := a + b
	 
	fmt.Println("The Addition is", c)
}

func mul() {

	e := 5
	f := 4
	d := e * f
	 
	fmt.Println("The multiplicaion is", d)
}


