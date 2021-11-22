package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/practice/bank_account/accounts"
)

func main() {
	minssogiAccount := accounts.CreateAccount("minssogi")

	fmt.Println("Deposit :", minssogiAccount.Deposit(50000))

	//fmt.Println(minssogiAccount)

	error := minssogiAccount.Withdraw(5000)
	if error == nil {
		fmt.Println("Withdraw :", minssogiAccount.Balance())
	}

	error = minssogiAccount.Withdraw(50000)
	if error == nil {
		fmt.Println("Withdraw :", minssogiAccount.Balance())
	} else {
		fmt.Println(error.Error())
	}

	//fmt.Println("Result:", minssogiAccount.Balance(), minssogiAccount.Name())
	fmt.Println("Result:", minssogiAccount)
}
