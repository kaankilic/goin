package main

import (
	"crypto/sha256"
	"time"
    "encoding/base64"
	"strconv"
	"fmt"
)

type Block struct {
	index int
	nonce int
	timestamp time.Time
	/* used */
	blockHash string
	pendingTransactions []Transaction
	previousHash string
}
func createBlock(pendingTransactions []Transaction, previousHash string) Block{
	blockHash := calculateHash(pendingTransactions, previousHash)
	return Block{blockHash: blockHash, previousHash:previousHash, pendingTransactions: pendingTransactions};
}
func calculateHash(pendingTransactions []Transaction, previousHash string) string{
	sha := sha256.New()
	sha.Write([]byte(time.Now().String()))
	sha.Write([]byte(fmt.Sprintf("%v", pendingTransactions)))
	sha.Write([]byte(previousHash))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil));
	return hash
}
func getTransactions(block Block) []Transaction{
	return block.transactions;
}