package main

import "fmt"

func main(){

	// using the append function we can add elements to the slice.

	var slic []int

	// declared the slice

	fmt.Println(slic) 
	// now the slice is empty so add elements using the append function.

	slic = append(slic,1,2)

	fmt.Println(slic)

	slic = append(slic,3,4,5,6,7,8)

	fmt.Println(slic)
}