package BlockChain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
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

func GenerateTransactionID(tx *Transaction) [32]byte {
	h := sha256.New()

	h.Write([]byte(tx.SenderAddress)) // Convert tx.SenderAddress to []byte

	binary.Write(h, binary.BigEndian, []byte(tx.SenderAddress))

	h.Write([]byte(tx.RecipientAddress))
	binary.Write(h, binary.BigEndian, tx.RecipientAddress)

	binary.Write(h, binary.BigEndian, tx.Amount)

	binary.Write(h, binary.BigEndian, tx.Time)

	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	h.Write(randomBytes)

	var ID [32]byte
	copy(ID[:], h.Sum(nil))

	return ID
}

func (w *Wallet) Serialise(Transaction *Transaction) []byte {
	data, _ := json.Marshal(*Transaction)
	return data
}

func (w *Wallet) MakeTransaction(RecipientAddress *rsa.PublicKey, Amount uint64) *Transaction {
	tx := NewTransaction(w.PublicKey, RecipientAddress, Amount)
	tx.Time = uint64(time.Now().UnixNano())
	ID := GenerateTransactionID(tx)
	tx.TransactionID = hex.EncodeToString(ID[:])
	Signature, err := SignTransaction(ID[:], w.PrivateKey) // Convert ID to []byte
	if err != nil {
		log.Fatal("Error signing")
	}
	tx.Signature = hex.EncodeToString(Signature)

	JsonTransaction := w.Serialise(tx)
	Client(&JsonTransaction)
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