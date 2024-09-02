package transactions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A txn is represented by a struct with sender's, reciever's Add and value sent
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// NewTransaction creates and returns a new Transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		senderBlockchainAddress:    sender,
		recipientBlockchainAddress: recipient,
		value:                      value,
	}
}

// GetSender returns the sender's blockchain address
func (t *Transaction) GetSender() string {
	return t.senderBlockchainAddress
}

// GetRecipient returns the recipient's blockchain address
func (t *Transaction) GetRecipient() string {
	return t.recipientBlockchainAddress
}

// GetValue returns the value of the transaction
func (t *Transaction) GetValue() float32 {
	return t.value
}

func (t *Transaction) PrintTxn() {
	fmt.Printf("%s\n", strings.Repeat("-", 60))
	fmt.Printf("sender_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value               %.1f\n", t.value)
}

// Marshal Txs struct to []byte
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
