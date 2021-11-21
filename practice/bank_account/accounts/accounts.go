package accounts

type Account struct {
	name    string
	balance int
}

func CreateAccount(name string) *Account {
	newAccount := Account{name: name}
	return &newAccount
}
