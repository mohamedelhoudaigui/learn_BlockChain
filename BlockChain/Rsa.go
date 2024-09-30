package BlockChain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func SignTransaction(TransactionByte []byte, PrivateKey *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(TransactionByte)
	signature, err := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hashed[:])
	return signature, err
}

func VerifyTransaction(TransactionByte []byte, Signature []byte, PublicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256(TransactionByte)
	return rsa.VerifyPKCS1v15(PublicKey, crypto.SHA256, hashed[:], Signature)
}
