package main

import (
	"fmt"
	"github.com/Mingsicogi/go_study/practice/bank_account/accounts"
)

func main() {
	minssogiAccount := accounts.CreateAccount("minssogi")

	fmt.Println(minssogiAccount.Deposit(50000))

	fmt.Println(minssogiAccount)
}
