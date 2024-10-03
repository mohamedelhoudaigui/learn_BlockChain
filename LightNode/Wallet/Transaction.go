package Wallet

import (
	"crypto/rsa"
	"fmt"
	"strings"
)

//------------impl-----------------

type	Transaction struct {
	TransactionID		[]byte `json:"transaction_id"`
	SenderAddress		*rsa.PublicKey `json:"sender_address"`
	RecipientAddress	*rsa.PublicKey `json:"recipient_address"`
	Amount				uint64 `json:"amount"`
	Time				uint64 `json:"time"`
	Signature			[]byte `json:"signature"`
}

func	(t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 60))
	fmt.Printf("sender address = %.10x...\n", t.SenderAddress)
	fmt.Printf("receiver address = %.10x...\n", t.RecipientAddress)
	fmt.Println("value =", t.Amount)
	fmt.Println("time =", t.Time)
	fmt.Printf("%s", strings.Repeat("*", 60))
}

func	NewTransaction(SenderAddress *rsa.PublicKey, RecipientAddress *rsa.PublicKey, Amount uint64) *Transaction {
	return &Transaction{
		SenderAddress:		SenderAddress,
		RecipientAddress:	RecipientAddress,
		Amount: 			Amount,
	}
}

//-------------------------------


