package models

import (
	json "encoding/json"
	"fmt"
	"strings"
)

type Transactions struct {
	senderBlockChainAddress   string
	receiverBlockChainAddress string
	amount                    float64
}

func NewTransaction(senderAddress, receiverAddress string, amount float64) *Transactions {
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
			Amount                    float64 `json:"amount"`
		}{
			SenderBlockChainAddress:   t.senderBlockChainAddress,
			ReceiverBlockChainAddress: t.receiverBlockChainAddress,
			Amount:                    t.amount,
		})
}
