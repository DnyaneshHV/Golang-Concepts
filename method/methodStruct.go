package main

import "fmt"

type Circle struct{

	radius float64
	area float64

}

func (c *Circle) calculateArea(){
	// If not pointer value of area will not get modified.

	const pi = 3.14

	c.area = pi*c.radius*c.radius

    
}

func main(){

	t := Circle{radius: 5}

	t.calculateArea()
	fmt.Println("Radius and Area",t)
}