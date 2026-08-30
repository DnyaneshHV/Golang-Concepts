package main

import "fmt"

type student struct {

	Name string
}

func (s student) disp(){

	fmt.Println("Name of the student is ", s.Name)
}

func main(){

	st := student{"John"}

	// store the address of st in p
	p := &st

	p.disp()   // Here go bahind the seen  , see this as (*p).disp()


	// Choosing the pointer reciver is good as we are not making any copies of the value to change them.

	
	
}