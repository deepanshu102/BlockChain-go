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

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}
func (b *Block) Print() {
	fmt.Printf("TimeStamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous Hash %s\n", b.previousHash)
	fmt.Printf("transactions %+v\n", b.transactions)
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
		Transactions []*Transactions `json:"transaction"`
	}{
		TimeStamp:    b.timestamp,
		Nonce:        b.nonce,
		Transactions: b.transactions,
		PreviousHash: b.previousHash,
	})
}
