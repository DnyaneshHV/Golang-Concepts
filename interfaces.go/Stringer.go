package main

import "fmt"

// Stringer is an inbuilt interface of Go language , used when we particularly want to print the 
// data in the string , or proper format 

// Declare the struct 

type Student struct{
	Name string
	Age int 
	Rollno int
}

//Now a function that implements the Stringer interface

func (s Student) String() string{

	// for printing the String method we use the Sprintf method of the fmt package

	return fmt.Sprintf(
		"User(Name=%s, Age=%d)",  // This is the format like %s represent string and %d represent integer.
		s.Name,
		s.Age,
	)
}



func main(){

	//Provide values to the struct

	 var p = Student{Name: "Dnyanesh", Age: 22}

	fmt.Println(p)



}

