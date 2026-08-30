package main

import "fmt"


type Person struct{

	Name string
	Age int
}

func (p Person) disp() {

	fmt.Println("Hello my name is ", p.Name)
}

func main(){

	// New topic methods begin.
	// Assign values to the struct
	person := Person{
		Name: "Dnyanesh",
		Age: 22,
	}

	person.disp()

}