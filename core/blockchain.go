package main

import (
	"crypto/sha256"
	"time"
    "encoding/base64"
	"strconv"
	"fmt"
)
type Blockchain struct{
	chain []Block
}
func createGenesisBlock() Block{
	return &Block{index:0,nonce:1,timestamp:"00/00/00",blockHash:"0",pendingTransactions:[],previousHash:nil}
}
func getLatestBlock(chain Blockchain) Block{
	return chain[len(chain)-1];
}
func addBlock(block Block){
	Chain = append(Chain, block)
}