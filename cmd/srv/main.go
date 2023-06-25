package main

import (
	"log"
	"git.deep.block/models"
)

func init(){
	log.SetPrefix("BlockChain:")
}
func main(){

	blockchain:=models.NewChain();
	blockchain.Print()
	previousHash:=blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5,previousHash)
	blockchain.Print()
	previousHash=blockchain.LastBlock().Hash()
	blockchain.CreateBlock(10, previousHash)
	blockchain.Print()
}