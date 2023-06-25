package models
import (
	"fmt"
	"strings"
)
type Chain struct{
	transactionsPool []string
	chains []*Block
}
func NewChain() *Chain{
	bc:=new(Chain)	
	bc.CreateBlock(0,"intial block")
	return bc
}
func (bc *Chain) CreateBlock(nonce int , previousHash string) *Block{
	block:=NewBlock(nonce, previousHash)
	bc.chains=append(bc.chains,block)
	return block
}
func (bc * Chain) Print(){
	for i, block:=range bc.chains{
		fmt.Printf("%s chain %d %s\n",strings.Repeat("=",25),i, strings.Repeat("=",25))
		block.Print()
	}
	fmt.Printf("%s\n",strings.Repeat("*",25))
}

func (bc *Chain) LastBlock() *Block{
	return bc.chains[len(bc.chains)-1]
}