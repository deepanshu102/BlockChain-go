package main

import (
	"log"

	"git.deep.block/models"
)

func init() {
	log.SetPrefix("BlockChain:")
}
func main() {

	blockchain := models.NewChain()
	blockchain.Print()
	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()
	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	previousHash = blockchain.LastBlock().Hash()
	nonce = blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()
}
