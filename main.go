package main

import (
	"toychain/blockchain"
)

// func MakeBlocks() {
// 	hashes := []string{"Genesis", "one", "two", "three", "four", "five"}
// 	for i := 0; i < 5; i++ {

// 		b := NewBlock(i, hashes[i])
// 		b.PrintBlock()
// 		println("****************************************************")
// 	}
// }

// func init() {
// 	log.SetPrefix("Blockchain: ")
// }

// func currentHash(hashIndx int) string {
// 	return fmt.Sprintf("hash:%d", hashIndx)
// }

func main() {
	/******************************  Nonce  ************************************************/
	blockChain := blockchain.NewBlockchain()
	// blockChain.PrintChain()

	blockChain.AddTransaction("A", "B", 1.2)
	blockChain.AddTransaction("C", "D", 2.4)
	blockChain.AddTransaction("E", "F", 2.3)

	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork() // Do POW to get nonce

	blockChain.AddBlock(nonce, previousHash)
	blockChain.PrintChain()

	/******************************  Transactions  ******************************************/

	// blockChain := blockchain.NewBlockchain()
	// blockChain.PrintChain() // only Genesis block in the chain
	// blockChain.PrintBlocks()

	// blockChain.LastBlock().PrintBlock()

	// blockChain.PrintTxnPool()
	// // All added Txn is in the TxnPool
	// blockChain.AddTransaction("A", "B", 1.2)
	// blockChain.AddTransaction("C", "D", 2.4)
	// blockChain.AddTransaction("E", "F", 2.3)
	// blockChain.AddTransaction("G", "H", 4.3)

	// blockChain.PrintTxnPool()

	// previousHash := blockChain.LastBlock().Hash()
	// blockChain.AddBlock(11111, previousHash)
	// // blockChain.PrintChain()

	// blockChain.AddTransaction("I", "J", 1.2)
	// blockChain.AddTransaction("K", "L", 2.4)
	// blockChain.AddTransaction("M", "N", 2.3)
	// blockChain.AddTransaction("O", "P", 4.3)
	// blockChain.AddTransaction("O", "R", 4.3)
	// blockChain.AddTransaction("O", "S", 4.3)

	// blockChain.PrintTxnPool()

	// blockChain.PrintChain()

	// blockChain.AddBlock(22222, previousHash)
	// blockChain.PrintTxnPool()
	// blockChain.PrintChain()
	/****************************** Hash *******************************************************/
	// myChain := blockchain.NewBlockchain()
	// // myChain.PrintChain()

	// previousHash := myChain.LastBlock().Hash()
	// myChain.AddBlock(50, previousHash)
	// myChain.LastBlock().PrintBlock()

	// println()

	// previousHash = myChain.LastBlock().Hash()
	// myChain.AddBlock(55, previousHash)
	// myChain.LastBlock().PrintBlock()

	// println()

	// previousHash = myChain.LastBlock().Hash()
	// myChain.AddBlock(75, previousHash)
	// myChain.LastBlock().PrintBlock()

	// myChain.PrintChain()

	/****************************** Blockchain Struct *******************************************************/
	// blockchain := blockchain.NewBlockchain()
	// blockchain.PrintChain()
	// for i := 0; i < 10; i++ {
	// 	blockchain.AddBlock(i, currentHash(i+1))
	// }
	// blockchain.PrintChain()

	// blockchain.AddBlock(0, "First")
	// blockchain.PrintChain()

	/******************************* Block Struct ***********************************************************/
	// println("****************************************************")
	// log.Println("Blockchain in Golang")
	// println("****************************************************")
	// println()
	// MakeBlocks()
}
