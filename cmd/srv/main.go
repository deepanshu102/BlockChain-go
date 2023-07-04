package main

import (
	"fmt"
	"log"

	"git.deep.block/block"
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

	// w := wallets.NewWallet()
	// fmt.Printf("%+v\n", w.PrivateKey())
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Printf("%+v\n", w.PublicKey())
	// fmt.Println(w.PublicKeyStr())
	// w := wallets.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())

	// t := wallets.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	// fmt.Printf("signature %s\n", t.GenerateSignature())
	walletM := wallets.NewWallet()
	walletA := wallets.NewWallet()
	walletB := wallets.NewWallet()

	// Wallet
	t := wallets.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Blockchain
	blockchain := block.NewChain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
