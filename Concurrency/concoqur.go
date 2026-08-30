package main 

import "fmt"
import "time"

// To achieve the go we use the go routine.

// a function that prins the particular word 5 times

func print(s string){

	for i :=0; i < len(s); i++ {
		// before printing we will stop it for some time
		time.Sleep(100 * time.Millisecond)

		// Can also use the wait groups , when we use them we dont have to write the sleep and all just use that packages or methods
		fmt.Println(s)
	}
}

func main(){

    go print("Hello")
	print("World")

}