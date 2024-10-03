package Wallet

import (
	"crypto/rsa"
	"fmt"
	"log"
	"strings"
	"time"
)


type	Wallet struct {
	PrivateKey	*rsa.PrivateKey
	PublicKey	*rsa.PublicKey
	Balance		uint64
}

func	(w *Wallet) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 60))
	fmt.Printf("Public key :	%v\n", w.PublicKey)
	fmt.Printf("Private key :	%v\n", w.PrivateKey)
	fmt.Printf("Balance : %d\n", w.Balance)
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

//-------------------------------



func (w *Wallet) MakeTransaction(RecipientAddress *rsa.PublicKey, Amount uint64, NodeAdress string) *Transaction {
	tx := NewTransaction(w.PublicKey, RecipientAddress, Amount)
	tx.Time = uint64(time.Now().UnixNano())
	ID := GenerateTransactionID(tx)
	tx.TransactionID = ID
	Signature, err := SignTransaction(w.PrivateKey, ID)
	if err != nil {
		log.Fatal("Error signing")
	}
	tx.Signature = Signature
	JsonTransaction := w.Serialise(tx)
	Client(&JsonTransaction, NodeAdress)
	return tx
}


//-------------------------

func	NewWallet() *Wallet {
	PrivateKey, err := GenerateKeyPair()

	if err != nil {
		log.Fatal("Error gen keys")
	}

	PublicKey := PrivateKey.PublicKey

	return &Wallet{
		PrivateKey:	PrivateKey,
		PublicKey:	&PublicKey,
		Balance:	0,
	}
}