package main

import (
	"fmt"
	"golang_til/appB/exB1/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
