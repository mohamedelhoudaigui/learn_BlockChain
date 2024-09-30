package BlockChain

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
	"log"
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

func GenerateTransactionID(tx *Transaction) [32]byte {
	h := sha256.New()

	h.Write(tx.SenderAddress.N.Bytes())
	binary.Write(h, binary.BigEndian, tx.SenderAddress.E)

	h.Write(tx.RecipientAddress.N.Bytes())
	binary.Write(h, binary.BigEndian, tx.RecipientAddress.E)

	binary.Write(h, binary.BigEndian, tx.Amount)

	binary.Write(h, binary.BigEndian, tx.Time.UnixNano())

	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	h.Write(randomBytes)

	var ID [32]byte
	copy(ID[:], h.Sum(nil))

	return ID
}

func	(w *Wallet) MakeTransaction(RecipientAddress *rsa.PublicKey, Amount uint64) *Transaction {
	tx := &Transaction{
		SenderAddress: w.PublicKey,
		RecipientAddress: RecipientAddress,
		Amount: Amount,
		Time: time.Now(),
	}

	ID := GenerateTransactionID(tx)
	tx.TransactionID = ID

	Signature, err := SignTransaction(tx.TransactionID[:], w.PrivateKey)
	if err != nil {
		log.Fatal("Error signing")
	}
	tx.Signature = Signature

	return tx
}


//-------------------------

func	NewWallet() *Wallet {
	PrivateKey, PublicKey, err := GenerateKeyPair()

	if err != nil {
		log.Fatal("Error gen keys")
	}

	return &Wallet{
		PrivateKey:	PrivateKey,
		PublicKey:	PublicKey,
		Balance:	0,
	}
}