package models

import (
	"fmt"
	"strings"
)

var (
	Mining_Difficulty = 3
)

type Chain struct {
	transactionsPool []*Transactions
	chains           []*Block
}

func NewChain() *Chain {
	block := &Block{}
	bc := new(Chain)
	bc.CreateBlock(0, block.Hash())
	return bc
}
func (bc *Chain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash)
	bc.chains = append(bc.chains, block)
	return block
}
func (bc *Chain) Print() {
	for i, block := range bc.chains {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Chain) LastBlock() *Block {
	return bc.chains[len(bc.chains)-1]
}
func (bc *Chain) AddTransaction(senderAddress, ReceiverAddress string, value float64) {
	transaction := NewTransaction(senderAddress, ReceiverAddress, value)
	bc.transactionsPool = append(bc.transactionsPool, transaction)
}
func (bc *Chain) CopyTransactionPool() []*Transactions {
	transactions := make([]*Transactions, 0)
	for _, t := range bc.transactionsPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockChainAddress, t.receiverBlockChainAddress, t.amount))
	}
	return transactions
}
func (bc *Chain) ValidPoof(nonce int, previousHash [32]byte, transaction []*Transactions, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{nonce, previousHash, 0, transaction}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	fmt.Println(guessHashStr)
	return guessHashStr[:difficulty] == zeros
}
func (bc *Chain) ProofOfWork() int {
	transaction := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidPoof(nonce, previousHash, transaction, Mining_Difficulty) {
		nonce += 1
	}
	return nonce
}
