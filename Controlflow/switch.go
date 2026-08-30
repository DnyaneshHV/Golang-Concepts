package main

import "fmt"

func main(){

	// This is the main function
	switchPractice()
	
}

func switchPractice(){

	// If the number matches then it is that day so 

	day := 6

	switch day {

	case 1:
		fmt.Println("Happy Monday")
	case 2:
		fmt.Println("Happy Tuesday")
	case 3:
		fmt.Println("Happy Wednesday")
	case 4:
		fmt.Println("Happy Thursday")
	case 5:
		fmt.Println("Happy Friday")
	case 6:
		fmt.Println("Happy Saturday")
	case 7:
		fmt.Println("Happy Sunday")
	default:
		fmt.Println("Invalid day")

    }
}