package BlockChain

import (
	"crypto/rsa"
	"fmt"
	"strings"
)

//------------impl-----------------

type MinerData struct {
	Chain	[]*Block
	Pool	[]*Transaction
	Diff	uint64
}

type	Transaction struct {
	TransactionID		[]byte			`json:"transaction_id"`
	SenderAddress		*rsa.PublicKey	`json:"sender_address"`
	RecipientAddress	*rsa.PublicKey	`json:"recipient_address"`
	Amount				uint64			`json:"amount"`
	Time				uint64			`json:"time"`
	Signature			[]byte			`json:"signature"`
}

func (t *Transaction) Print() {
    fmt.Println(strings.Repeat("-", 50))
    fmt.Printf("%-15s: %s\n", "Transaction ID", truncateAndFormat(t.TransactionID))
    fmt.Printf("%-15s: %s\n", "Sender", truncateAndFormat(t.SenderAddress.N.Bytes()))
    fmt.Printf("%-15s: %s\n", "Recipient", truncateAndFormat(t.RecipientAddress.N.Bytes()))
    fmt.Printf("%-15s: %d\n", "Amount", t.Amount)
    fmt.Printf("%-15s: %s\n", "Time", formatTime(t.Time))
    fmt.Printf("%-15s: %s\n", "Signature", truncateAndFormat(t.Signature))
    fmt.Println(strings.Repeat("-", 50))
}

//-------------------------------

func	NewTransaction(SenderAddress *rsa.PublicKey, RecipientAddress *rsa.PublicKey, Amount uint64) *Transaction {
	return &Transaction{
		SenderAddress:		SenderAddress,
		RecipientAddress:	RecipientAddress,
		Amount: 			Amount,
	}
}

func	NewMinerData(bc *BlockChain) *MinerData {
	return &MinerData{
		Chain: bc.Chain,
		Pool: bc.TransactionPool,
		Diff: bc.Diffic,
	}
}



