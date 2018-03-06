package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

type Block struct {
	index     int
	nonce     int
	timestamp time.Time
	/* used */
	blockHash           string
	pendingTransactions []Transaction
	previousHash        string
}

func createBlock(index int, nonce int, timestamp time.Time, pendingTransactions []Transaction, previousHash string) Block {
	blockHash := calculateHash(index, nonce, timestamp, pendingTransactions, previousHash)
	return Block{blockHash: blockHash, previousHash: previousHash, pendingTransactions: pendingTransactions}
}
func calculateHash(index int, nonce int, timestamp time.Time, pendingTransactions []Transaction, previousHash string) string {
	sha := sha256.New()
	sha.Write([]byte(time.Now().String()))
	sha.Write([]byte(fmt.Sprintf("%v", pendingTransactions)))
	sha.Write([]byte(previousHash))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return hash
}
