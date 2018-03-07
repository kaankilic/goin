package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

type Block struct {
	index               int
	nonce               int
	timestamp           time.Time
	blockHash           string
	pendingTransactions []Transaction
	previousHash        string
}

func createBlock(index int, nonce int, timestamp time.Time, pendingTransactions []Transaction) Block {
	return Block{index: index, nonce: nonce, timestamp: timestamp, pendingTransactions: pendingTransactions}
}
func (self *Block) setPreviousHash(previousHash string) {
	self.previousHash = previousHash
}
func (self *Block) setBlockHash(blockHash string) {
	self.blockHash = blockHash
}
func (self *Block) calculateHash() string {
	fmt.Println("+++++++++++++++++++++")
	fmt.Println("calc hash timestamp:" + self.timestamp.String())
	fmt.Println("calc hash pendingTransactions:" + fmt.Sprintf("%v", self.pendingTransactions))
	fmt.Println("calc hash previousHash:" + self.previousHash)
	fmt.Println("calc hash nonce:" + string(self.nonce))
	fmt.Println("calc hash index:" + string(self.index))
	fmt.Println("+++++++++++++++++++++")
	sha := sha256.New()
	sha.Write([]byte(self.timestamp.String()))
	sha.Write([]byte(fmt.Sprintf("%v", self.pendingTransactions)))
	sha.Write([]byte(string(self.nonce)))
	sha.Write([]byte(string(self.index)))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return hash
}
