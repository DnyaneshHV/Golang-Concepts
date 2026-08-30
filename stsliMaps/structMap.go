package main

import "fmt"

func main(){

	// we can also use struct in the map

	type s = struct{

		x , y int
	}

	// x age and y roll no

	info := map[string]s{
		"Ram":{18,11},
		"Ganesh":{54,12},
		"Viraj":{75,15},
	}

	fmt.Println(info)

	updateMap()
}

//Muatating the map means insert or update te map .

func updateMap(){

	var city = map[string]int{

		"Pune":410501,
		"Nagpur":540958,
		"Nashik":534878,
		"Mumbai":393823,
	}

	fmt.Println(city)
	city["Pune"] = 49333

	fmt.Println(city)

	// Delete with the key

	delete (city,"Nagpur")  // This is how we delete.

	fmt.Println()
	fmt.Println(city)
}