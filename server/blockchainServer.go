package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"git.deep.block/block"
	"git.deep.block/utils"
	"git.deep.block/wallets"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

type BlockChainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockChainServer {
	return &BlockChainServer{port}
}
func (bcs *BlockChainServer) Port() uint16 {
	return bcs.port
}
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world")
}

func (bcs *BlockChainServer) GetBlockChain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallets.NewWallet()
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("Private key  %v", minerWallet.PrivateKeyStr())
		log.Printf("Public key  %v", minerWallet.PublicKeyStr())
		log.Printf("block-Blockchain Address  %v", minerWallet.BlockchainAddress())

	}
	return bc
}

func (bcs *BlockChainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockChain()
		m, err := bc.MarshalJSON()
		if err != nil {
			log.Print(err)
		}
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: INVAID METHOD")
	}
}

func (bcs *BlockChainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockChain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			Transactions []*block.Transaction `json:"transactions"`
			Length       int                  `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m[:]))

	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			fmt.Printf("\n\n%+v\n\n\n", t)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		fmt.Printf("Public key:- %064x%064x\n", publicKey.X.Bytes(), publicKey.Y.Bytes())
		bc := bcs.GetBlockChain()
		isCreated := bc.CreateTransaction(*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, *t.Value, publicKey, signature)

		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockChainServer) Mine(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bc := bcs.GetBlockChain()
		isMined := bc.Mining()
		var m []byte
		if !isMined {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusOK)
			m = utils.JsonStatus("success")
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error:Invalid HTTP Method")
	}
}
func (bcs *BlockChainServer) Run() {
	http.HandleFunc("/chains", bcs.GetChain)
	http.HandleFunc("/transactions", bcs.Transactions)
	http.HandleFunc("/mine", bcs.Mine)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
