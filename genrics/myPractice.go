package main

import "fmt"


func Add[T int | float32] (a, b T) T{

	// meaning of above syntax is T can be int or float32 and variable a and b 
	// can be of type T and the function returns the value  in the type T
	return a + b
}

func main(){

	fmt.Println("The addition of two numbers is : ", Add(12,13))
}