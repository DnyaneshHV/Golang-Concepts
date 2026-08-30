package main

import "fmt"


func f1(a int, ch chan int){

	// send the value received in the channel , this will be the bidirectionla channel

	ch <- a
}

func main(){

	c := make(chan int)

	fmt.Println("The channle and the Goroutines example")

	go f1(10,c)

	n := <-c

	fmt.Println("Receiving the value of n: ", n)

}


// we are sending and receiving the values between the main and the another function goroution
// This is how we use the channels