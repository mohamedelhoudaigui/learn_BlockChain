package BlockChain

import (
	"crypto/rsa"
	"fmt"
	"strings"
)

//------------impl-----------------

type	Transaction struct {
	TransactionID		string `json:"transaction_id"`
	SenderAddress		string `json:"sender_address"`
	RecipientAddress	string `json:"recipient_address"`
	Amount				uint64 `json:"amount"`
	Time				uint64 `json:"time"`
	Signature			string `json:"signature"`
}

func	(t *Transaction) Print() {
	fmt.Printf("%s", strings.Repeat("=", 25))
	fmt.Println("sender address =", t.SenderAddress)
	fmt.Println("receiver address =", t.RecipientAddress)
	fmt.Println("value =", t.Amount)
	fmt.Println("time =", t.Time)
	fmt.Printf("%s", strings.Repeat("*", 60))
}

func	NewTransaction(SenderAddress *rsa.PublicKey, RecipientAddress *rsa.PublicKey, Amount uint64) *Transaction {
	return &Transaction{
		SenderAddress: SenderAddress.N.String(),
		RecipientAddress: RecipientAddress.N.String(),
		Amount: Amount,
	}
}

//-------------------------------


