package main

import "fmt"

func main(){

	// Map declaration using the literals and using make.

	ages := map[string]int{

		"John":25,
		"Ram":26,
	}

	fmt.Println("Map elements: ")
	fmt.Println(ages)



	fmt.Println()
	fmt.Println("Another way to declare maps is as follows: ")

	// we  use make keyword

	city := make(map[int]string)

	city[540] = "Pune"
	city[539] = "Nashik"
	city[938] = "Mumbai"

	fmt.Println(city)

	fmt.Println()
	mutat()

}

// Map mutate means changing of the map elements like delete, remove , change.

func mutat(){

	// slice declare

	student := []string{"ram","sham","ganesh","viraj","kishore"}

	fmt.Println(student)

	india := make(map[string]int)  // declaration of the map

	india["maharashtra"] = 1  //Insert
	india["gujrat"] = 2
	india["keralam"] = 3
	india["Goa"] = 4

	fmt.Println(india)

	fmt.Println()  // delete keralam

	delete(india,"keralam")
    fmt.Println(india)

	//now update the map value
	india["Goa"] = 101
	fmt.Println(india)

	value, ok := india["gujrat"]
    
	if ok {
		fmt.Println("gujrat exists with the value ", value)
	}else{
		fmt.Println("gujrat does not exitst")
	}
}