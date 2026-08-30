package main

import "fmt"

// In other languages ws have classes and when we write functions in those classes we call 
// as methods

// Here in Go language when 
// WE dont have classes , we have structs and we need to bring our functions to this structs 
// And then they will be called as methods 

type Person struct{
	Name string,
	Email string,
	Status  bool,
	Age int
}

func main(){

	hitesh:= Person{Name: Hitesh, Email: hitesh@gmail.com, Status: true, Age: 23}

	hitesh.Getstatus()

}

// Create the method to see the status of the person 

func (p Person) Getstatus(){
 
	fmt.Println("The staus of the person is", p.Status)
}