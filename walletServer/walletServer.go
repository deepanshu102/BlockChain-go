package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"

	"git.deep.block/block"
	"git.deep.block/utils"
	"git.deep.block/wallets"
)

const tempDir = "walletServer/templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{
		port:    port,
		gateway: gateway,
	}
}
func (ws *WalletServer) Port() uint16 {
	return ws.port
}
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Print("	Error: INVALID METHOD")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallets.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		decorder := json.NewDecoder(r.Body)
		var t wallets.TransactionRequest
		err := decorder.Decode(&t)
		if err != nil {
			log.Printf("ERROR:%v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Printf("ERROR:Fields are missing")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		// fmt.Println(*t.RecipientBlockchainAddress)
		// fmt.Println(*t.SenderBlockchainAddress)
		// fmt.Println(*t.SenderPrivateKey)
		// fmt.Println(*t.SenderPublicKey)
		// fmt.Println(*t.Value)

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		value64 := float32(value)
		fmt.Println(publicKey)
		fmt.Println(privateKey)
		fmt.Println(value64)

		WalletTransaction := wallets.NewTransaction(privateKey, publicKey, *t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, value64)
		signature := WalletTransaction.GenerateSignature()
		signatureStr := signature.String()
		blockChainTransaction := &block.TransactionRequest{
			SenderBlockchainAddress:    t.SenderBlockchainAddress,
			RecipientBlockchainAddress: t.RecipientBlockchainAddress,
			Value:                      &value64,
			SenderPublicKey:            t.SenderPublicKey,
			Signature:                  &signatureStr,
		}
		m, _ := json.Marshal(blockChainTransaction)
		fmt.Printf("%+v\n", string(m[:]))
		buf := bytes.NewBuffer(m)
		resp, err := http.Post(ws.Gateway()+"/transactions", "application/json", buf)
		if err != nil {
			log.Printf("ERROR:%v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}
		io.WriteString(w, string(utils.JsonStatus("fail")))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Amount(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		blockchainAddress := r.URL.Query().Get("blockchain_address")
		endpoint := fmt.Sprintf("%s/amount", ws.Gateway())
		blockchainServerRquest, _ := http.NewRequest(http.MethodGet, endpoint, nil)
		q := blockchainServerRquest.URL.Query()
		q.Add("blockchain_address", blockchainAddress)
		blockchainServerRquest.URL.RawQuery = q.Encode()
		bcsResponse, err := http.DefaultClient.Do(blockchainServerRquest)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		if bcsResponse.StatusCode == 200 {
			decoder := json.NewDecoder(bcsResponse.Body)
			var bar block.AmountResponse
			err := decoder.Decode(&bar)
			if err != nil {

				log.Printf("ERROR: %v", err)
				io.WriteString(w, string(utils.JsonStatus("fails")))
				return
			}
			m, _ := json.Marshal(struct {
				Message string  `json:"message"`
				Amount  float32 `json:"amount"`
			}{
				Message: "success",
				Amount:  bar.Amount,
			})
			io.WriteString(w, string(m[:]))
			return
		} else {
			io.WriteString(w, string(utils.JsonStatus("fails")))
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/wallet/amount", ws.Amount)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))

}
