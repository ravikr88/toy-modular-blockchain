package main

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
	/****************************** Transactions ************************************************/
	// blockChain := blockchain.NewBlockchain()
	// blockChain.PrintChain() // only Genesis block in the chain

	// // All added Txn is in the TxnPool
	// blockChain.AddTransaction("A", "B", 1.2)
	// blockChain.AddTransaction("C", "D", 2.4)
	// blockChain.AddTransaction("E", "F", 2.3)

	// previousHash := blockChain.LastBlock().Hash() // Genesis block
	// blockChain.AddBlock(11111, previousHash)
	// blockChain.PrintChain()

	// // some more new txns created after first block is created, will be added in next block
	// blockChain.AddTransaction("CCC", "DDD", 2.0)
	// blockChain.AddTransaction("XXX", "YYY", 3.0)
	// blockChain.AddBlock(11111, previousHash)

	/****************************** Hash *******************************************************/
	// blockchain := blockchain.NewBlockchain()
	// blockchain.PrintChain()

	// previousHash := blockchain.LastBlock().Hash()
	// blockchain.AddBlock(1231323, previousHash)
	// blockchain.PrintChain()

	// previousHash = blockchain.LastBlock().Hash()
	// blockchain.AddBlock(321322, previousHash)
	// blockchain.PrintChain()
	//
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
