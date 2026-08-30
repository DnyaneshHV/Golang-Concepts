package main

import (
	"fmt"
	"time"
)

func main() {


	ti := time.Now()

	switch {

		case ti.Hour() < 12: 
		  fmt.Println("Good Morning")
	    case ti.Hour() < 17:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening")

	}

}