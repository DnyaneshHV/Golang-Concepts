package main

import "fmt"

type Calculation interface{
	 AreaBoth() float64  // using one single method we will be able to use for 2 or more structs
}

// Above declared the interface 

// Type will implement the interface so lets make another type

type square struct{
	// Area of the square is 4 multiply by side

	Area float64   // go not allow like same field and the method name
	side float64
}

type rect struct {

	Area float64
	length float64
	breadth float64
}

func (s square) AreaBoth() float64{

	// 

	s.Area = 4 * s.side

	return s.Area

}

func (r rect) AreaBoth() float64{

	// interfaces are implemented implicitly we dont need to implement them externally or explicitly

	// rectangleArea method implemented by the type rect , this is example of interface 
	// And any number of types can use those methods no problem

	r.Area = r.length * r.breadth

	return r.Area

}

// func printData(s square){
// 	fmt.Println("Field values - ",s)
// 	s.AreaBoth()
// 	fmt.Println("Area of Square", s)  // after updating the value s will be shown
// }

// // This practise is not good like for writing the same code for data of square and rectangle to print not good.

// func printDataRec(r rect){
// 	fmt.Println("Field Values-",r)
// 	// inside r we have area and side
// 	fmt.Println("Area of Rectangle", r.AreaBoth())
// 	// in method we require variables that we take from here 
// }

func print(c Calculation){

	fmt.Println("The Values : ",c)
	fmt.Println("The Area : ",c.AreaBoth())


}

func main(){
	fmt.Println("Area of Square and Rectangle is as follows....")

	sqr := square{side: 5}
	rec := rect{length: 4, breadth: 7}

	print(sqr)
	print(rec)   // when we use the pointer we have to give its address.


	// Empty interface is called as nil interface.

	// var i interface{}

	// The underlying type of the variable does not changes 
	// But in case of interface the underlying type changes at the run time

	// can see using "%T\n" , this gives the type
	
}