package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//EncrptAES takes in a data slice and a passphrase and returns a aes encrypted string
func EncryptAES(data []byte, passphrase string) []byte {

	//Creating a aes cipher block with a md5 passphrase hash
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	//passes block into the Galios counter measure
	gcm, _ := cipher.NewGCM(block)
	//Allocate a byte sized for the gcm nonce
	nonce := make([]byte, gcm.NonceSize())
	//Create a nonce to satisfy encryption
	io.ReadFull(rand.Reader, nonce)
	//Acual encryption
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func DecryptAES(data []byte, passphrase string) []byte {
	key := []byte(CreateHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	decrypted, _ := gcm.Open(nil, nonce, cipherText, nil)
	return decrypted
}

func FileEncrypt(f string) {

}
