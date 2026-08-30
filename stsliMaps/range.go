package main

import "fmt"

func main(){

	// Range for loop gives us two things the index and the value.

	var s = []int{1,2,3,4,5,6,7,8}

	// iterate wite the for loop of the range

	for i,v := range s{

		fmt.Println("The index : ", i ,"The value : ",v)
	}

	contRange()
}

// Range continue

func contRange(){

	var sli = []string{"ram","sham","ganesh","suraj","vishal","gaurav"}


	fmt.Println("The one by one values of the slice is as follows: ")
	for _,v := range sli {

		// we have to keep the space for the index i
		
		fmt.Println(v)
	}
}