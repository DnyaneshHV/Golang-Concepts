package main 

import "fmt"


type name struct{
	Name string
}

func main(){

	var i interface{}
	var j interface{}

    j = name{Name: "Deer"}

	i = 10

	s, ok := i.(int)  // we are telling go like get me value of type int from the interface

	fmt.Println(s)
	fmt.Println(ok)

	sj := j.(name)
	 
	fmt.Println(j)
	fmt.Println(sj)

	fmt.Printf("%T\n", j)
    fmt.Printf("%T\n", sj)
}

