package wallets

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockChainAddress string
}

/*
use for create a New Wallet
*/
func NewWallet() *Wallet {
	w := new(Wallet)
	// Step 1: create ECDSA Private key [ 32 bytes] .. Public Key [64 Bytes]
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	// Step 2: Perform SHA-256 hashing on public key[32 bytes]
	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// Setp 3: Perfom RIPEMD-160 hashing on the result of SHA-256 [20 BYTES]
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// step 4:
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// step 5:
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// step 6:
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// step 7:
	checksum := digest6[:4]
	// step 8:
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], checksum[:])
	// step 9:
	address := base58.Encode(dc8)
	w.blockChainAddress = address
	return w
}

// Get blockchain address
func (w *Wallet) BlockchainAddress() string {
	return w.blockChainAddress
}

// PrivateKey() is
// use for get access of private key of particular wallet
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// PrivateKeyStr() is help to get Private key in readable format
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

// PublicKey() is used for get details of PublicKey of particular wallet
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// PublicKeyStr() is hepl to get Public key in readable format
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%064x%064x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"private_key"`
		PublicKey         string `json:"public_key"`
		BlockchainAddress string `json:"blockchain_address"`
	}{
		PrivateKey:        w.PrivateKeyStr(),
		PublicKey:         w.PublicKeyStr(),
		BlockchainAddress: w.BlockchainAddress(),
	})
}
