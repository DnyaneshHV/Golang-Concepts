package oprations

import (
	"fmt"

)


func DepositeMoney(accountNumber int, password string, amount float64){

	// check if the account number and password is correct or not
	
	for i, u := range Users{
		if u.AccountNumber == accountNumber && u.Password == password{
			// if the account number and password is correct then deposit the amount in his account
			Users[i].Balance += amount
			fmt.Println("Amount Deposited Successfully")
			fmt.Println("Current Balance: ", Users[i].Balance)

			fmt.Println("Amount Deposited Successfully")
			return
		}
	}
	
	fmt.Println("Invalid Account Number or Password")
	return
}

