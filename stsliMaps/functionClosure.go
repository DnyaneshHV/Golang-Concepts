package main

import "fmt"

func aggro() func(int) int {

	sum := 0
	
	return func(x int) int{
		sum += x
		return sum
	}
}

// store that function in the main

func main(){

	post := aggro()

	fmt.Println(post(1))
	fmt.Println(post(2))
	fmt.Println(post(3))

}