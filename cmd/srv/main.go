package main

import (
	"fmt"
	"log"

	"git.deep.block/wallets"
)

func init() {
	log.SetPrefix("BlockChain:")
}
func main() {
	// myBlockChainAddress := "myBlockChainAddress"
	// blockChain := models.NewChain(myBlockChainAddress)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 1.0)
	// previousHash := blockChain.LastBlock().Hash()
	// nonce := blockChain.ProofOfWork()
	// blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Mining()
	// blockChain.Print()

	// blockChain.AddTransaction("C", "D", 2.0)
	// blockChain.AddTransaction("X", "Y", 10.0)
	// blockChain.Mining()
	// previousHash = blockChain.LastBlock().Hash()
	// nonce = blockChain.ProofOfWork()
	// blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 1.0)
	// blockChain.Mining()
	// blockChain.Print()

	// blockChain.AddTransaction("C", "D", 2.0)
	// blockChain.AddTransaction("X", "Y", 3.0)
	// blockChain.Mining()
	// blockChain.Print()
	// fmt.Printf("my %.1f\n", blockChain.CalculatTotalAmount(myBlockChainAddress))
	// fmt.Printf("C %.1f\n", blockChain.CalculatTotalAmount("C"))
	// fmt.Printf("D %.1f\n", blockChain.CalculatTotalAmount("D"))

	w := wallets.NewWallet()
	fmt.Printf("%+v\n", w.PrivateKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Printf("%+v\n", w.PublicKey())
	fmt.Println(w.PublicKeyStr())
}
