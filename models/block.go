package models

import(
	"time"
	"fmt"
	json "encoding/json"
)
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
	fmt.Printf("TimeStamp %d\n",b.timestamp)
	fmt.Printf("nonce %d\n",b.nonce)
	fmt.Printf("previous Hash %s\n",b.previousHash)
	fmt.Printf("transactions %s\n",b.transactions)
}

func (b *Block)Hash() string{
	m,_:=json.Marshal(b)
	return string(m)
}

func (b *Block) MarshalJSON()([]byte,error){
	return json.Marshal(struct{
		TimeStamp int64 `json:"timestamp"`
		Nonce int `json:"nonce"`
		PreviousHash string `json:"previousHash"`
		Transactions []string `json:"transaction"`
	}{
		TimeStamp:b.timestamp,
		Nonce:b.nonce,
		Transactions:b.transactions,
		PreviousHash:b.previousHash,
	})
}