package main

import (
	"fmt"

	"github.com/tlee0818/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Taco")
	account.Deposit(10)
	fmt.Println(account.getBalance())

}
