package main

import "fmt"

func main(){

	// slice literals are like array without the size

	// int bool struct

	var q = []int{1,2,3,4,5,6,7,8,9,10}

	var p = []bool{true,false,false,true,false,true}

	fmt.Println(q)
	fmt.Println(p)

	// also we can create a struct

	var st = []struct{
		i int
		y string
	}{
		{1,"Ram"},
		{2,"Sham"},
		{3,"Ganesh"},
		{4,"Shubhash"},
		{5,"Amar"},
		{6,"kishore"},

	}

	fmt.Println(st)

	defaultslice()
}

func defaultslice(){

	// Default slice is like what happen when we dont give starting or ending point to the slice for like underlaying array.

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)   // each slice operation is perform based on the result of the previous slice.

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}