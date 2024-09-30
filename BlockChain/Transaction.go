package BlockChain

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"
)

//------------impl-----------------

type	Transaction struct {
	TransactionID		[32]byte
	SenderAddress		*rsa.PublicKey
	RecipientAddress	*rsa.PublicKey
	Amount				uint64
	Time				time.Time
	Signature			[]byte
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
		SenderAddress: SenderAddress,
		RecipientAddress: RecipientAddress,
		Amount: Amount,
	}
}

//-------------------------------


