package main

import (
	"fmt"
	"strings"
)

// here the concept of maps , loops and string will be used

func main(){

	var example string = "go is fun go is fast"

// have to count the maximum times the word appear

// string.Fields that splits the string into the words

words := strings.Fields(example)

// in words the words are stored in the form of the slice

//make the map

freq := make(map[string]int) // this creates the empty map

 for _, word := range words {
	freq[word]++
	// above statement will increase the frequency of the go statements
 }

 // now the values are stored in the map and iterate through the map and print values

 for word, count := range freq{
	fmt.Println(word, count)
 }

}




