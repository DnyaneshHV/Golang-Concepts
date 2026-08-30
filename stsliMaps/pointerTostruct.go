package main

import "fmt"

type Person struct {

	Name string
	Age *int   // we have the address of the Age that should point to the value
}

// Can also do like as follows

type Display struct{

	x, y int
}

var(

	v= Display{1,2}
	v1= Display{x:1}
	f= &Display{1,2}
)

func main() {

	a := 9

	// Have to give values to the struct
	p := Person {
		Name: "Ganesh",
		Age : &a,
	}

	fmt.Println("The Age points to the a: ", *p.Age) // without the star it will print the address of where a location in the memory.


	//--------------------------------------------------  Can do also like this.
     fmt.Println(v1)
	 fmt.Println(v)
	 fmt.Println(f)

}