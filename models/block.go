package models

// Its one block struct and that represent the block of our Blockchain
type Block struct{
	nonce int 
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string)*Block{
	b:=new(Block)
	b.timestamp=time.Now().UnixNano()
	b.nonce=nonce
	b.previousHash=previousHash
	return b
}
func (b *Block) Print(){
	fmt.Printf("TimeStamp %d\n",b.time)
	fmt.Printf("nonce %d\n",b.nonce)
	fmt.Printf("previous Hash %s\n",b.previousHash)
	fmt.Printf("transactions %s\n",b.transactions)
}


