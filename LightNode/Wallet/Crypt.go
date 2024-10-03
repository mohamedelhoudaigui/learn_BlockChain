package Wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
)

func GenerateKeyPair() (*rsa.PrivateKey, error) {
    return rsa.GenerateKey(rand.Reader, 2048)
}


func SignTransaction(privateKey *rsa.PrivateKey, ID []byte) ([]byte, error) {
    hashed := sha256.Sum256(ID)
    return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
}

func GenerateTransactionID(tx *Transaction) []byte {
	TxJson, _ := json.Marshal(*tx)
	hash := sha256.Sum256(TxJson)
	return hash[:]
}

func (w *Wallet) Serialise(tx *Transaction) []byte {
	data, _ := json.Marshal(*tx)
	return data
}
