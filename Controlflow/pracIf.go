package main

import "fmt"



func main(){
 
	var a int

	fmt.Println("Enter age to check your elegiblity for voting: ")
	fmt.Scan(&a)

	if a > 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	disp()
}

// if with the short statement prints the 1 to 5 numbers

func disp() {

	if value := 1; value < 6 {
		fmt.Println(value)
	}
	
}