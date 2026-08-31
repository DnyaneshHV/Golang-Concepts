package main

import "fmt"

func main(){

	var light = []int{1,2,3,4,5,6,7}
	max := light[0]  // given the first value.
	min := light[0]


	// want to find the largest and the lowest value in the slice

	for _, num := range light{

		if num > max{
			max = num
		}

		if num < min {
			min = num
		}

	}

	fmt.Println("MaxValue is : ", max)
	fmt.Println("MinValue is : ", min)

}