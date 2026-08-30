package oprations

import "fmt"

func CheckBalance(accountNumber int, password string){

	// check if the account number and password is correct or not
	
	for _, u := range Users{
		if u.AccountNumber == accountNumber && u.Password == password{
			// if the account number and password is correct then display the balance of his account
			fmt.Println("Current Balance: ", u.Balance)
			return
		}
	}
	
	fmt.Println("Invalid Account Number or Password")
	return
}