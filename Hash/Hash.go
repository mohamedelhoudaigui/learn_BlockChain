package Hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetHash(chainN int, time int) string {
	data := fmt.Sprintf("%d%d", chainN, time)
	hash := sha256.New()
	hash.Write([]byte(data))
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}
