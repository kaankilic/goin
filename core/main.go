package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	bc := createBlockchain()
	bc.createGenesisBlock()
	transactions := []Transaction{}
	w1 := createWallet()
	w2 := createWallet()
	fmt.Println("wallet 1:" + w1.walletHash)
	fmt.Println("wallet 1 Balance:" + strconv.Itoa(w1.balance))
	fmt.Println("wallet 2:" + w2.walletHash)
	fmt.Println("wallet 2 Balance:" + strconv.Itoa(w2.balance))
	w1.decreaseAmount(10)
	transactions = append(transactions, createTransaction(w1, w2, 10))
	w2.decreaseAmount(5)
	transactions = append(transactions, createTransaction(w2, w1, 5))
	w1.decreaseAmount(7)
	transactions = append(transactions, createTransaction(w1, w2, 7))
	fmt.Println("transactions length:" + strconv.Itoa(len(transactions)))
	block := createBlock(1, 1, time.Now(), transactions)
	bc.addBlock(block)
	bc.addBlock(block)
	bc.addBlock(block)
	bc.addBlock(block)
	fmt.Println("w1 balance:" + strconv.Itoa(w1.balance))
	fmt.Println("w2 balance:" + strconv.Itoa(w2.balance))
	fmt.Println("is chain valid:" + strconv.FormatBool(bc.isValid()))
}
