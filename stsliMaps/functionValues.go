package main

import "fmt"

func main(){

	// Functions can also be passed as the values.

	//declare the variable that can store any function with two variables and that return one integer value.

	var operation func(int , int) int  // declared function that can store this type of value

	// store the add funcion inside the variable

	operation = add  // stores the add function inside the variable.

	result := operation(10,20)

	fmt.Println(result)

}

func add(a,b int) int {

	return a + b
}