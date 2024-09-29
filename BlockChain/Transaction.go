package BlockChain

import (
	"fmt"
	"strings"
	"time"
)

type	Transaction struct {
	SenderAddress		string
	RecipientAddress	string
	Value				float64
	Time				time.Time
}

func	(t *Transaction) Print() {
	fmt.Printf("%s", strings.Repeat("=", 25))
	fmt.Printf("sender address = %s\n", t.SenderAddress)
	fmt.Printf("receiver address = %s\n", t.RecipientAddress)
	fmt.Printf("value = %.1f\n", t.Value)
	fmt.Printf("time = %v\n", t.Time)
	fmt.Printf("%s", strings.Repeat("*", 60))
}

//-----------------------------

func	NewTransaction(SenderAddress string, RecipientAddress string, Value float64) *Transaction {
	return &Transaction{
		SenderAddress: SenderAddress,
		RecipientAddress: RecipientAddress,
		Value: Value,
	}
}

