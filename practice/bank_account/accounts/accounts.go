package accounts

import "github.com/Mingsicogi/go_study/practice/bank_account/common_constants"

type Account struct {
	name    string
	balance int
}

func CreateAccount(name string) *Account {
	newAccount := Account{name: name}
	return &newAccount
}

func (account Account) Balance() int {
	return account.balance
}

// method(receiver)
func (account *Account) Deposit(amount int) int {
	account.balance += amount
	return account.balance
}

// error handling
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return common_constants.ErrInsufficientBalance
	}

	account.balance -= amount
	return nil
}
