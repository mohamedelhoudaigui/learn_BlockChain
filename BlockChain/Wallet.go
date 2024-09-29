package BlockChain

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"log"
	"time"
)


type	Wallet struct {
	PrivateKey	[32]byte
	PublicKey	[32]byte
	Balance		uint64
}

func	GenKeys() ([32]byte, [32]byte) {
	r := rand.Reader
	seed, err := rand.Int(r, new(big.Int).SetInt64(272727272727272727))
	if err != nil {
		log.Fatal(err) // or return an error
	}
	t := time.Now().UnixNano()
	PublicKey := sha256.Sum256([]byte(fmt.Sprintf("%d%d", seed, t)))
	PrivateKey := sha256.Sum256([]byte(fmt.Sprintf("%d%d%x", seed, t, PublicKey)))
	return PublicKey, PrivateKey
}

func	(w *Wallet) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 60))
	fmt.Printf("Public key :	%x\n", w.PublicKey)
	fmt.Printf("Private key :	%x\n", w.PrivateKey)
	fmt.Printf("Balance : %d\n", w.Balance)
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func	(w *Wallet) MakeTransaction(RecipientAddress [32]byte, Amount uint64) *Transaction {
	tx := &Transaction{
		SenderAddress: w.PublicKey,
		RecipientAddress: RecipientAddress,
		Amount: Amount,
		Time: time.Now(),
	}
	Id := fmt.Sprint("%x%x%d%d", tx.SenderAddress, tx.RecipientAddress, tx.Amount, time.Now().UnixNano())
	tx.TransactionID = sha256.Sum256([]byte(Id))
	var Signature [32]byte
	for i:= 0; i < 32; i++ {
		Signature[i] = (tx.TransactionID[i] + w.PrivateKey[i]) % 255
	}
	tx.Signature = Signature
	return tx
}


//-------------------------

func	NewWallet() *Wallet {
	PrivateKey, PublicKey := GenKeys()
	return &Wallet{
		PrivateKey:	PrivateKey,
		PublicKey:	PublicKey,
		Balance:	0,
	}
}