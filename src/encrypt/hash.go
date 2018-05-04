package encrypt

import (
	"crypto/sha512"
	"encoding/hex"
)

func createHash(key string) string {
	hasher := sha512.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
