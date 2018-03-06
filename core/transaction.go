package main

type Transaction struct {
	from   Wallet
	to     Wallet
	amount int
}

func createTransaction(sender Wallet, receiver Wallet, amount int) Transaction {
	if sender.balance < amount {
		receiver.balance = sender.balance - amount
		sender.balance = receiver.balance + amount
		return Transaction{from: sender, to: receiver, amount: amount}
	}
	return Transaction{}
}
