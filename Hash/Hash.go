package Hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetHash(chainN int, time int, prHash string) string {
	data := fmt.Sprintf("%d%d%s", chainN, time, prHash)
	hash := sha256.New()
	hash.Write([]byte(data))
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}
