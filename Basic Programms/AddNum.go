package main

import "fmt"

func addNum(a int, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 5

	result := addNum(a, b)
	fmt.Println("The sum of a and b is :", result)
}