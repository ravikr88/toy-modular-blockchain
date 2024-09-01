package blockchain

import (
	"fmt"
	"strings"
	b "toychain/block"
)

// A Blockchain is represented by a struct with 2 attributes
type Blockchain struct {
	transactionPool []string   // all Txn created on chain, to be mined and added to chain
	chain           []*b.Block // array of Blocks
}

// NewBlockchain creates a new chain
func NewBlockchain() *Blockchain {
	b := &b.Block{}

	// new returns pointer to the created Blockchain
	bc := new(Blockchain)

	// Genesis Block for new chain
	bc.AddBlock(0, b.Hash())

	// return the new chain
	return bc
}

// AddBlock method on a chain creates a new block, add it to chain, and returns a pointer to that added block
func (bc *Blockchain) AddBlock(nonce int, previousHash [32]byte /*string*/) *b.Block {
	// create a new Block
	b := b.NewBlock(nonce, previousHash)

	// add the new block to []*Block or chain
	bc.chain = append(bc.chain, b)

	// return the added block
	return b
}

// LastBlock returns the latest / Last block added in the chain
func (bc *Blockchain) LastBlock() *b.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) PrintChain() {
	// Iterate over all blocks in the chain and print each block
	for i, block := range bc.chain {
		if i == 0 {
			fmt.Printf("%s %s  %s\n", strings.Repeat("=", 25), "GENESIS BLOCK", strings.Repeat("=", 25))
		} else {
			fmt.Printf("%s Block Number: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		}
		block.PrintBlock()
	}
	fmt.Printf("%s %s %s\n", strings.Repeat("*", 25), "end of chain", strings.Repeat("*", 25))
	fmt.Println()
}
