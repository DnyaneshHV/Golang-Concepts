package main

import "fmt"

func main(){

	// Take input from the user and store them in array and slice.
 

	// declare the slice
	var n int
	arr := make([]int,0,n)
	
	fmt.Println("How many numbers do you want to enter: ")
	fmt.Scanln(&n)

	fmt.Println("Enter the numbers : ")

	for i :=0; i < n ;i++{

		var x int 
		fmt.Scan(&x)

		arr = append(arr, x)
		
	}

	fmt.Println("Slice : ", arr)
	usingArray()
}


func usingArray(){

	// declare

	const n = 5
	var arr1 [n]int 

	fmt.Println("Enter the numbers: ")

	for i:= 0 ; i < n; i++{

		fmt.Scan(&arr1[i])
	}
	fmt.Println(arr1)
}
	