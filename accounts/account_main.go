package main

import (
	"fmt"
	"log"

	"github.com/Hyuk-II/Go_basic/accounts"
)


func main() {
	account := accounts.NewAccount("hyuk")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account)

	err := account.WithDraw(10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
}