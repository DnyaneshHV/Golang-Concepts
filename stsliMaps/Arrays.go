package main

import "fmt"

func main(){

	//Array declaration

	var number [5]int 

	number = [5]int{1,2,3,4,5}

	// when used := it means create new and assign, but only = is aggisn this values.

	fmt.Println(number)
	loops()
	fmt.Println()
	Str()
}

func loops(){

	// loop throught the array

	var arr [5]int
	arr = [5]int{10,20,30,40,50}

	// iterate through the array and print the values.

	for i:=0; i < len(arr); i++{
		fmt.Println(arr[i])
	}
}

func Str(){

	var st [2]string
	st = [2]string{"Ram","Sham"}

	fmt.Println(st[0])
	fmt.Println(st[1])
}