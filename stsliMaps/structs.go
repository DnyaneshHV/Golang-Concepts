package main

import "fmt"

func main(){

	// structs are like storing the  related data values in the one field

	type Person struct{
		Name string
		Age int
		city string
		mob int32 
	}

	// we saved the related types of the field in one struct

	// now use it

	person := Person{Name: "Gaesh" , Age : 25 , city : "Pune" , mob : 435678343}

	fmt.Println("The Information of the person is: ", person)

	fmt.Println("Access using the dot: ", person.Name)
	fmt.Println("Access using the dot: ", person.Age)
	fmt.Println("Access using the dot: ", person.city)


}

// We can accss the struct fields using the . operator.