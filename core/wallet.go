package main

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

type Wallet struct {
	walletHash string
	balance    int
}

func createWallet() Wallet {
	sha := sha256.New()
	sha.Write([]byte(time.Now().String()))
	return Wallet{walletHash: base64.URLEncoding.EncodeToString(sha.Sum(nil)), balance: 100}
}
func (self *Wallet) decreaseAmount(amount int) {
	self.balance -= amount
}
func (self *Wallet) increaseAmount(amount int) {
	self.balance += amount
}
