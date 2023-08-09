package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"git.deep.block/block"
	"git.deep.block/wallets"
)

var cache map[string]*block.Chain = make(map[string]*block.Chain)

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

func (bcs *BlockChainServer) GetBlockChain() *block.Chain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallets.NewWallet()
		bc = block.NewChain(minerWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("Private key  %v", minerWallet.PrivateKeyStr())
		log.Printf("Public key  %v", minerWallet.PublicKeyStr())
		log.Printf("block-chain Address  %v", minerWallet.BlockchainAddress())

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

func (bcs *BlockChainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
