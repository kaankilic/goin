package main

import (
	"fmt"
	"strconv"
)

func main() {
	transactions := []Transaction{}
	w1 := createWallet()
	w2 := createWallet()
	fmt.Println("wallet 1:" + w1.walletHash)
	fmt.Println("wallet 1 Balance:" + strconv.Itoa(w1.balance))
	fmt.Println("wallet 2:" + w2.walletHash)
	fmt.Println("wallet 2 Balance:" + strconv.Itoa(w2.balance))
	transactions = append(transactions, createTransaction(w1, w2, 10))
	transactions = append(transactions, createTransaction(w2, w1, 5))
	transactions = append(transactions, createTransaction(w1, w2, 7))
	block := createBlock(transactions, "0")
	fmt.Println(block)
	fmt.Println("w1 balance:" + strconv.Itoa(w1.balance))
	fmt.Println("w2 balance:" + strconv.Itoa(w2.balance))
}
