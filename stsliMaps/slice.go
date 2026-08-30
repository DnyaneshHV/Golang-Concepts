package main

import "fmt"

func main(){

	// slice is underlying growing array we dont need to specify the size of an slice

	slice := []int{1,2,3,4,5,6,7,8}

	for i:=0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// In slice we not mention the size.

	// we can create the slice of an existing array.
	sli()
}

func sli(){

	var arr [5]int 

	arr = [5]int{11,22,33,44,55}

	// now printing the sile from the above array

	var sli []int = arr[1:4]
	fmt.Println("The silce or the array is : ", sli)

	// In slice when we make chages for like any value it becomes permanent.
	
	b :=  arr[1:4]

	fmt.Println(b)

	b[0] = 80

	fmt.Println(b)
	fmt.Println(arr) // value in the array also got changed.
}

