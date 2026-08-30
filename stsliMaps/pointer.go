package main

import "fmt"

func main() {

	// Pointer is something that like in which we store the memory address of the value stored.

	// We use & for like to get address of variable and we use * for like change the value of that particular variable 

	p := 10

	// we have to store the address of p in particular variable 

	x := &p

	// we want to print the address of p

	fmt.Println("Address of p: ", x)

	// Now want to change the value of p using the pointer

	// Printing the value of x
	fmt.Println("Value of p ", *x)

	// Change x so p will be changed automatically

	*x = 20  // value changed

	fmt.Println("Changed value is : ", p)
}