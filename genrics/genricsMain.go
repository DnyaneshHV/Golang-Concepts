package main

import "fmt"

// T is a placeholder for any type
func Firstitems []T T {
	return items[0]
}

func main() {
	// T becomes int
	numbers := []int{10, 20, 30}
	fmt.Println("First Number:", First(numbers))

	// T becomes string
	names := []string{"John", "David", "Tom"}
	fmt.Println("First Name:", First(names))

	// T becomes float64
	prices := []float64{99.99, 49.50, 25.00}
	fmt.Println("First Price:", First(prices))
}