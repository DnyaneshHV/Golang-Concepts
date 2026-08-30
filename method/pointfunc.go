package main

import "fmt"

// relation between pointer and function

type name struct{
	name string
	age int
}

func main(){

	info := name{"ganesh",24}

	disp(&info)   // This &info gives me the memory address of the info not the values.
   	fmt.Println(info)


}

func disp(n *name){   // but the * what does it do is go to the pointer and access the actual values.

	n.age = 25

	// behind the seen go does is *n.age , so this concept is called as automatic pointer indirection
	

}