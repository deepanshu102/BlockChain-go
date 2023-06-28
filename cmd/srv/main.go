package main

import (
	"log"

	"git.deep.block/models"
)

func init() {
	log.SetPrefix("BlockChain:")
}
func main() {
	// myBlockChainAddress := "myBlockChainAddress"
	blockChain := models.NewChain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()
}
