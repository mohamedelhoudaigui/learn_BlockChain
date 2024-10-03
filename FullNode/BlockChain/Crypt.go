package BlockChain

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"time"
)

// Verify signature
func VerifyTransaction(publicKey *rsa.PublicKey, ID, signature []byte) error {
    hashed := sha256.Sum256(ID)
    return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

func parsePublicKey(publicKeyPEM string) (*rsa.PublicKey, error) { // move from string to rsa.PublicKey
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("key type is not RSA")
	}
}

func truncateAndFormat(data []byte) string {
    if len(data) == 0 {
        return "N/A"
    }
    hexStr := hex.EncodeToString(data)
    if len(hexStr) > 10 {
        return hexStr[:10] + "..."
    }
    return hexStr
}

func formatTime(timestamp uint64) string {
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}