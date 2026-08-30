package main

import "fmt"

type name []string // this is the type of the slice


func (n name) print(){

	// using the for loop print those names
	for i, name := range n{
		fmt.Println(i, name)
	}
}


func main(){

	friend := name{"Dan","Marie","John","Ganesh"}

	friend.print()
}