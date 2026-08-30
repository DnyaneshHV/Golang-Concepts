package main

import (
	"fmt"
	"MenuDrivnPro/oprations"
)

func main() {

	// In go the for loop acts as the do , while and for loop
	
    // Bank management system in golang

	for{
		var choice int
		// --------------- Make the menu

		fmt.Println("Welcome to the Bank Management System")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Check Balance")
		fmt.Println("5. View Account Details")
		fmt.Println("6. Exit")
		
		//------- Take the value and store in the variable declared above.
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
 

		// use the switch case to perform the particular operation based on the users input

		switch choice{

			case 1:
				user := oprations.UserAccount{}
				user.CreateAccountUser()
			case 2:
				var accountNumberD int
				var passwordD string
				var amount float64
				fmt.Println("Enter your Account Number: ")
				fmt.Scanln(&accountNumberD)
				fmt.Println("Enter your Password: ")
				fmt.Scanln(&passwordD)
				fmt.Println("Enter the Amount to be Deposited: ")
				fmt.Scanln(&amount)
				oprations.DepositeMoney(accountNumberD, passwordD, amount)
			case 3: 
			    var withDrawAccountNumber int
				var withDrawPassword string
				var withDrawAmount float64
				fmt.Println("Enter your Account Number: ")
				fmt.Scanln(&withDrawAccountNumber)
				fmt.Println("Enter your Password: ")
				fmt.Scanln(&withDrawPassword)
				fmt.Println("Enter the Amount to be Withdrawn: ")
				fmt.Scanln(&withDrawAmount)
				oprations.WithDrawMoney(withDrawAccountNumber, withDrawPassword, withDrawAmount)	


			case 4:
				var accountNumberB int
				var passwordB string
				fmt.Println("Enter your Account Number: ")
				fmt.Scanln(&accountNumberB)
				fmt.Println("Enter your Password: ")
				fmt.Scanln(&passwordB)
				oprations.CheckBalance(accountNumberB, passwordB)

			case 5:
				//view account details
				// here we just have to take account number and password.
				var accountNumber int
				var password string
				fmt.Println("Enter your Account Number: ")
				fmt.Scanln(&accountNumber)
				fmt.Println("Enter your Password: ")
				fmt.Scanln(&password)
				oprations.ViewAccountDetails(accountNumber, password)
				// Taken values and passed to the funcion.

			case 6:
				fmt.Println("Exiting from the bamk")
				fmt.Println("Thank you for visiting the KATARI BANK")
				return 
			default :
			    fmt.Println("Invalid user input")

		}
	}

}