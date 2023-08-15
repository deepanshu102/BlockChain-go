package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"git.deep.block/utils"
)

var (
	MINING_DIFFICULTY         = 3
	MINING_SENDER             = "Miner Address"
	MINING_REWARD     float32 = 1.0
)

// Its one block struct and that represent the block of our Blockchain
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transactions
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transactions) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}
func (b *Block) Print() {
	fmt.Printf("TimeStamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous Hash %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TimeStamp    int64           `json:"timestamp"`
		Nonce        int             `json:"nonce"`
		PreviousHash string          `json:"previousHash"`
		Transaction  []*Transactions `json:"transactions"`
	}{
		TimeStamp:    b.timestamp,
		Nonce:        b.nonce,
		Transaction:  b.transactions,
		PreviousHash: fmt.Sprintf("%x", b.previousHash),
	})
}

type Transactions struct {
	senderBlockChainAddress   string
	receiverBlockChainAddress string
	amount                    float32
}

func NewTransaction(senderAddress, receiverAddress string, amount float32) *Transactions {
	return &Transactions{
		senderBlockChainAddress:   senderAddress,
		receiverBlockChainAddress: receiverAddress,
		amount:                    amount,
	}

}

func (t *Transactions) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("Sender Address: %s\n", t.senderBlockChainAddress)
	fmt.Printf("Receiver Address %s\n", t.receiverBlockChainAddress)
	fmt.Printf("Value : %.1f\n", t.amount)
}

func (t *Transactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			SenderBlockChainAddress   string  `json:"senderBlockChainAddress"`
			ReceiverBlockChainAddress string  `json:"receiverBlockchainAddress"`
			Amount                    float32 `json:"amount"`
		}{
			SenderBlockChainAddress:   t.senderBlockChainAddress,
			ReceiverBlockChainAddress: t.receiverBlockChainAddress,
			Amount:                    t.amount,
		})
}

type Chain struct {
	transactionsPool  []*Transactions
	chains            []*Block
	blockChainAddress string
	port              uint16
}

func NewChain(blockChainAddress string, port uint16) *Chain {
	block := &Block{}
	bc := new(Chain)
	bc.port = port
	bc.blockChainAddress = blockChainAddress
	bc.CreateBlock(0, block.Hash())
	return bc
}

func (bc *Chain) TransactionPool() []*Transactions {
	return bc.transactionsPool
}
func (bc *Chain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Block []*Block `json:"chains"`
	}{
		Block: bc.chains,
	})
}

func (bc *Chain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	block := NewBlock(nonce, previousHash, bc.transactionsPool)
	bc.chains = append(bc.chains, block)
	bc.transactionsPool = []*Transactions{}
	return block
}
func (bc *Chain) LastBlock() *Block {
	return bc.chains[len(bc.chains)-1]
}
func (bc *Chain) Print() {
	for i, block := range bc.chains {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}
func (bc *Chain) CreateTransaction(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature) bool {
	isTransacted := bc.AddTransaction(sender, recipient, value, senderPublicKey, s)

	// TODO
	// Sync

	return isTransacted
}

func (bc *Chain) AddTransaction(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == MINING_SENDER {
		bc.transactionsPool = append(bc.transactionsPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		/*
			if bc.CalculateTotalAmount(sender) < value {
				log.Println("ERROR: Not enough balance in a wallet")
				return false
			}
		*/
		bc.transactionsPool = append(bc.transactionsPool, t)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}
	return false

}

func (bc *Chain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transactions) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
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
	for !bc.ValidPoof(nonce, previousHash, transaction, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Chain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockChainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action:mining, Status=success")
	return true
}

func (bc *Chain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chains {
		for _, t := range b.transactions {
			value := t.amount
			if blockchainAddress == t.receiverBlockChainAddress {
				totalAmount += value
			}
			if blockchainAddress == t.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

type TransactionRequest struct {
	SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
	Value                      *float32 `json:"value"`
	SenderPublicKey            *string  `json:"sender_public_key"`
	Signature                  *string  `json:"signature"`
}

func (tr *TransactionRequest) Validate() bool {
	if tr.RecipientBlockchainAddress == nil || tr.Signature == nil || tr.Value == nil || tr.SenderBlockchainAddress == nil || tr.SenderPublicKey == nil {
		return false
	}
	return true
}
