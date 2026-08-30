package oprations

import "fmt"

func WithDrawMoney(accountNumber int, password string, amount float64){

	// check if the account number and password is correct or not
	
	for i, u := range Users{
		if u.AccountNumber == accountNumber && u.Password == password{
			// if the account number and password is correct then withdraw the amount from his account
			if u.Balance >= amount{
				Users[i].Balance -= amount
				fmt.Println("Amount Withdrawn Successfully")
				fmt.Println("Current Balance: ", Users[i].Balance)
				return
			}else{
				fmt.Println("Insufficient Balance")
				return
			}
		}
	}
	
	fmt.Println("Invalid Account Number or Password")
	return
}