package models

import (
	"crypto/sha256"
	json "encoding/json"
	"fmt"
	"time"
)

// Its one block struct and that represent the block of our Blockchain
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transactions
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transactions) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}
func (b *Block) Print() {
	fmt.Printf("TimeStamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous Hash %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TimeStamp    int64           `json:"timestamp"`
		Nonce        int             `json:"nonce"`
		PreviousHash [32]byte        `json:"previousHash"`
		Transaction  []*Transactions `json:"transactions"`
	}{
		TimeStamp:    b.timestamp,
		Nonce:        b.nonce,
		Transaction:  b.transactions,
		PreviousHash: b.previousHash,
	})
}
