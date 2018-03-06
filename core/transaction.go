package main

import (
	"fmt"
	"strconv"
)

type Transaction struct {
	from   Wallet
	to     Wallet
	amount int
}

func createTransaction(sender Wallet, receiver Wallet, amount int) Transaction {
	if sender.balance > amount {
		fmt.Println("sender wallet:" + sender.walletHash)
		fmt.Println("receiver wallet:" + receiver.walletHash)
		fmt.Println("amount:" + strconv.Itoa(amount))
		return Transaction{from: sender, to: receiver, amount: amount}
	}

	return Transaction{}
}
