package oprations

import "fmt"

// declare the struct for like the variables of the user account creation

var Users []UserAccount

type UserAccount struct{
	Name string
	Email string
	Password string
	PhoneNumber string
	AccountNumber int
	Balance float64

}

func (u *UserAccount)CreateAccountUser(){
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&u.Name)
	fmt.Println("Enter your Email: ")
	fmt.Scanln(&u.Email)
	fmt.Println("Enter your Password: ")
	fmt.Scanln(&u.Password)
	fmt.Println("Enter your Phone Number: ")
	fmt.Scanln(&u.PhoneNumber)
	fmt.Println("Enter your Account Number: ")
	fmt.Scanln(&u.AccountNumber)

	Users = append(Users, *u)

	fmt.Println("Account Created Successfully")

	return

}


// View Account Details 

func ViewAccountDetails(accountNumber int, password string){
	// user passed form main.go is stored in u of type UserAccount

	for _, u := range Users{
		if u.AccountNumber == accountNumber && u.Password == password{
			fmt.Println("Account Details: ")
			fmt.Println("Name: ", u.Name)
			fmt.Println("Email: ", u.Email)
			fmt.Println("Phone Number: ", u.PhoneNumber)
			fmt.Println("Account Number: ", u.AccountNumber)
		
			return
		}
	}
	
	fmt.Println("Invalid Account Number or Password")
	return

}