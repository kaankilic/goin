package main

import (
	"time"
)

var chain []Block

func createGenesisBlock() *Block {
	now := time.Now()
	return &Block{index: 0, nonce: 1, timestamp: now, blockHash: "0", pendingTransactions: nil, previousHash: "0x"}
}
func getLatestBlock() Block {
	return chain[len(chain)-1]
}
func addBlock(block Block) {
	chain = append(chain, block)
}
