package accounts

type Account struct {
	name    string
	balance int
}

func CreateAccount(name string) *Account {
	newAccount := Account{name: name}
	return &newAccount
}

// method(receiver)
func (account *Account) Deposit(amount int) int {
	account.balance += amount
	return account.balance
}
