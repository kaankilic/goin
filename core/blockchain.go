package main

import (
	"fmt"
	"time"
)

type Blockchain struct {
	chain []Block
}

func createBlockchain() Blockchain {
	return Blockchain{chain: nil}
}
func (self *Blockchain) createGenesisBlock() {
	now := time.Now()
	gBlock := Block{index: 0, nonce: 1, timestamp: now, blockHash: "0", pendingTransactions: nil, previousHash: "0x"}
	self.chain = append(self.chain, gBlock)
}
func (self *Blockchain) getLatestBlock() Block {
	return self.chain[len(self.chain)-1]
}

func (self *Blockchain) addBlock(block Block) {
	block.setBlockHash(block.calculateHash())
	block.setPreviousHash(self.getLatestBlock().blockHash)
	self.chain = append(self.chain, block)
}
func (self *Blockchain) isValid() bool {
	for i := 1; i < len(self.chain); i++ {
		currentBlock := self.chain[i]
		previousBlock := self.chain[i-1]
		fmt.Println("current Block hash:" + currentBlock.blockHash)
		fmt.Println("current calculated Block hash:" + currentBlock.calculateHash())
		if currentBlock.blockHash != currentBlock.calculateHash() {
			return false
		}
		if currentBlock.previousHash != previousBlock.blockHash {
			return false
		}
	}
	return true
}
