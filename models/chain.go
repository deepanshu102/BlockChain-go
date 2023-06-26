package models

import (
	"fmt"
	"strings"
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
