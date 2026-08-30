package main

import "fmt"

// Generic function
func Printvaluevalue T {
	fmt.Println(value)
}

func main() {
	Printvalue(10)      // int
	Printvalue("Hello") // string
	Printvalue(true)    // bool
}