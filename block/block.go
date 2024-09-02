package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
	txn "toychain/transactions"
)

// A block is represented by a struct with 4 attributes
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*txn.Transaction
}

// Setter methods to initialize private fields
func (b *Block) SetNonce(nonce int) {
	b.nonce = nonce
}

func (b *Block) SetPreviousHash(previousHash [32]byte) {
	b.previousHash = previousHash
}

func (b *Block) SetTransactions(transactions []*txn.Transaction) {
	b.transactions = transactions
}

func (b *Block) SetTimestamp(timestamp int64) {
	b.timestamp = timestamp
}

// Getter methods to access private fields (if needed)
func (b *Block) GetNonce() int {
	return b.nonce
}

func (b *Block) GetPreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) GetTransactions() []*txn.Transaction {
	return b.transactions
}

func (b *Block) GetTimestamp() int64 {
	return b.timestamp
}

// NewBlock() takes nonce generated by performing POW, txns from Pool  and previousHash and creates a new block
func NewBlock(nonce int, previousHash [32]byte /*string*/, transactions []*txn.Transaction) *Block {
	// new returns a pointer to a new Block strcut
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// Print the entire block
func (b *Block) PrintBlock() {
	fmt.Printf("timestamp      	%d\n", b.timestamp)
	fmt.Printf("nonce          	%d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)

	// print all txn in the block
	for _, t := range b.transactions {
		t.PrintTxn()
	}

}

// Custom marshaling function
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64              `json:"timestamp"`
		Nonce        int                `json:"nonce"`
		PreviousHash [32]byte           `json:"previous_hash"`
		Transactions []*txn.Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// Hash a block and return a 256-bit output or 32 bytes
func (b *Block) Hash() [32]byte {
	m, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(m)
}
