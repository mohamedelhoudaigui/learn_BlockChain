package BlockChain

import (
	"fmt"
	"strings"
	"time"
)

type	Transaction struct {
	TransactionID		[32]byte
	SenderAddress		[32]byte
	RecipientAddress	[32]byte
	Amount				uint64
	Time				time.Time
	Signature			[32]byte
}

func	(t *Transaction) Print() {
	fmt.Printf("%s", strings.Repeat("=", 25))
	fmt.Printf("sender address = %s\n", t.SenderAddress)
	fmt.Printf("receiver address = %s\n", t.RecipientAddress)
	fmt.Printf("value = %.1f\n", t.Amount)
	fmt.Printf("time = %v\n", t.Time)
	fmt.Printf("%s", strings.Repeat("*", 60))
}

//-----------------------------



//-----------------------------

func	NewTransaction(SenderAddress string, RecipientAddress string, Amount float64) *Transaction {
	return &Transaction{
		SenderAddress: SenderAddress,
		RecipientAddress: RecipientAddress,
		Amount: Amount,
	}
}
