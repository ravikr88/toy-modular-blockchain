package blockchain

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	b "toychain/block"
	txn "toychain/transactions"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

// A Blockchain is represented by a struct with 2 attributes
type Blockchain struct {
	transactionPool   []*txn.Transaction // all Txn created on chain, to be mined and added to chain
	chain             []*b.Block         // chain is an array of Blocks
	blockchainAddress string
}

// func (bc *Blockchain) PrintEntireChain() {
// 	fmt.Println(bc.chain[0])
// 	fmt.Println(bc.transactionPool[0])
// }

// NewBlockchain creates a new chain
func NewBlockchain(blockchainAddress string) *Blockchain {
	// empty block
	b := &b.Block{}

	// new returns pointer to the empty crated Blockchain
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress

	// Genesis Block for new chain using empty block's hash
	bc.AddBlock(0, b.Hash())

	// return the new chain
	return bc
}

// AddBlock method on a chain creates a new block, add it to chain, and returns a pointer to that added block
func (bc *Blockchain) AddBlock(nonce int, previousHash [32]byte /*string*/) *b.Block {
	// create a new Block
	b := b.NewBlock(nonce, previousHash, bc.transactionPool)

	// add the new block to []*Block or chain
	bc.chain = append(bc.chain, b)

	// empty the txn pool after coppying all txns
	bc.transactionPool = []*txn.Transaction{}

	// return the added block
	return b
}

// LastBlock returns the latest / Last block added in the chain
func (bc *Blockchain) LastBlock() *b.Block {
	return bc.chain[len(bc.chain)-1]
}
func (bc *Blockchain) PrintTxnPool() {
	for i, txn := range bc.transactionPool {
		fmt.Printf("txn no: %d %v\n", i, *txn)
		// fmt.Println(*txn)

	}
}
func (bc *Blockchain) PrintBlocks() {
	for i, block := range bc.chain {
		fmt.Printf("block no: %d %+v\n", i, block)
	}
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
	// for i, block := range bc.chain {
	// 	fmt.Printf("%s Chain Height %d %s\n", strings.Repeat("=", 25), i,
	// 		strings.Repeat("=", 25))
	// 	block.PrintBlock()
	// }
	// fmt.Printf("%s %s %s\n", strings.Repeat("*", 25), "end of chain", strings.Repeat("*", 25))
	// fmt.Println()
}

// AddTransaction creates a new Txn and appends it to TxnPool in the chain
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := txn.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) PrintCopiedTxnPool() {
	copiedTxns := bc.CopyTnxPool()
	for _, txns := range copiedTxns {
		fmt.Printf("%v\n", *txns)
	}
}

// CoppyTxnPool returns an array of copied Txns from TxnPool
func (bc *Blockchain) CopyTnxPool() (copiedTxns []*txn.Transaction) {
	copiedTxns = make([]*txn.Transaction, 0)

	// Iterate over transactions stored in the transaction pool
	for _, t := range bc.transactionPool {
		// Create a new transaction by copying the details of the existing one using getter methods
		newTxn := txn.NewTransaction(t.GetSender(), t.GetRecipient(), t.GetValue())
		copiedTxns = append(copiedTxns, newTxn)
		// fmt.Println(t.GetRecipient())
	}

	return copiedTxns
}

// ValidProof checks if the hash of the block has the required number of leading zeros
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, txns []*txn.Transaction, DIFFICULTY int) bool {
	// Required zeros in hash eg. 0000 for DIFFICULTY 4
	zeros := strings.Repeat("0", DIFFICULTY)

	// Create a new empty guess block
	// all fields of guess block is respective zero value
	guessBlock := b.Block{}

	// Initialize the guess block using setter methods
	guessBlock.SetNonce(nonce)
	guessBlock.SetPreviousHash(previousHash)
	guessBlock.SetTransactions(txns)
	guessBlock.SetTimestamp(0) //  can set an actual timestamp if needed

	// Hash the guess block
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	fmt.Printf("%x", guessHashStr)
	// Convert the hash to a hex string

	// Return true if the initial zeros of the hash match the required zeros
	return guessHashStr[:DIFFICULTY] == zeros
}
func (bc *Blockchain) ProofOfWork() (nonce int) {
	txns := bc.CopyTnxPool()
	previousHash := bc.LastBlock().Hash()

	nonce = 0

	for {
		// Create a new guess block
		guessBlock := b.Block{}
		guessBlock.SetNonce(nonce)
		guessBlock.SetPreviousHash(previousHash)
		guessBlock.SetTransactions(txns)
		guessBlock.SetTimestamp(0) // Can set an actual timestamp if needed

		// Hash the guess block
		guessHash := guessBlock.Hash()
		guessHashHex := hex.EncodeToString(guessHash[:])

		// Print nonce and hash for debugging
		// fmt.Printf("Trying nonce %d: Computed hash: %s\n", nonce, guessHashHex)

		// Check if the computed hash has the required number of leading zeros
		zeros := strings.Repeat("0", MINING_DIFFICULTY)
		if guessHashHex[:MINING_DIFFICULTY] == zeros {
			// fmt.Printf("Found valid nonce %d: Hash with required zeros: %s\n", nonce, guessHashHex)
			break
		}

		// Increment nonce
		nonce++
	}

	// Return the nonce that satisfies the proof of work
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.AddBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}
