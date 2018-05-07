package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

//CreateHash takes in a string and returns a sha512 hashed string
func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
