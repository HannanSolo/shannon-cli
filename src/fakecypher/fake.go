package fakecypher

import "crypto/cipher"

type fakeCipher struct{}

func NewCipher(key []byte) (cipher.Block, error) { return fakeCipher{}, nil }

func (fakeCipher) BlockSize() int { return 16 }

func (fakeCipher s) Encrypt(dst, src []byte) {
	copy(dst, src[:s.BlockSize()])
}

func (fakeCipher s) Decrypt(dst, src []byte) {
	copy(dst, src[:s.BlockSize()])
}
