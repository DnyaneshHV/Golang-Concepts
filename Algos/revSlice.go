package main

import "fmt"

func main(){

	// reverse the slice

	var sli []int

	sli = []int{1,2,3,4,5,6,7,8}

	// Slice dont have the builtin reverse method

	// can create a function for slice , for reversing the slice
	reverse(sli)

	fmt.Println(sli)

	
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}