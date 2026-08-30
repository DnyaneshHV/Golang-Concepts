package main

import "fmt"

// User implements the error interface
type User struct {
	Name string
	Age  int
}

// Error must RETURN a string
func (u User) Error() string {
	return fmt.Sprintf(
		"%s is not eligible for voting because age is %d",
		u.Name,
		u.Age,
	)
}

func checkVotingEligibility(name string, age int) error {

	// Error scenario
	if age < 18 {
		return User{
			Name: name,
			Age:  age,
		}
	}

	// No error
	return nil
}

func main() {

	var name string
	var age int

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	err := checkVotingEligibility(name, age)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("You are eligible for voting")
}