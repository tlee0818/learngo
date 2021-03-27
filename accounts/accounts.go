package accounts

// account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates account
// doesnt return a copy, but the actual object itself (the address)
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0}

	return &account
}

//deposit x amount to the account
func (a Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) getBalance() int {
	return a.balance
}
