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

func createBlock(index int, nonce int, pendingTransactions []Transaction) Block {
	timestamp := time.Now()
	return Block{index: index, nonce: nonce, timestamp: timestamp, pendingTransactions: pendingTransactions}
}
func (self *Block) setPreviousHash(previousHash string) {
	self.previousHash = previousHash
}
func (self *Block) setBlockHash(blockHash string) {
	self.blockHash = blockHash
}
func (self *Block) calculateHash() string {
	sha := sha256.New()
	sha.Write([]byte(string(self.timestamp.Unix())))
	sha.Write([]byte(fmt.Sprintf("%v", self.pendingTransactions)))
	sha.Write([]byte(string(self.nonce)))
	sha.Write([]byte(string(self.index)))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return hash
}
