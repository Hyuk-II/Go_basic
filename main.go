package main

import (
	"fmt"

	"github.com/Hyuk-II/Go_basic/accounts"
)


func main() {
	account := accounts.NewAccount("hyuk")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())
}