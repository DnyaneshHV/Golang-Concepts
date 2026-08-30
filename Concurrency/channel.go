package main

import "fmt"

func main(){

	// declaration of the channel 

	var ch chan int
	// int type of data will be sent through the channel

	// initialization of the channel
	fmt.Println(ch)
	
	ch = make(chan int)

	// now 
	fmt.Println(ch)

	// Channels are alwasy used with Goroutines as they are used to communicate between differnt go routines.
	

}