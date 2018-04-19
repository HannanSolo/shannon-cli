package main_test

import (
	"crypto/aes"
	"fmt"
	"testing"
)

// this test should really have been broken down into several tests
// each test can be used to check your assumptions about how things work
// when writing tests for new code, you can use tests to assert your assumptions about how it *should* work
func TestBlockInterface(t *testing.T) {
	key := make([]byte, 16)
	cipher, err := aes.NewCipher(key) // TODO replace this with FakeCipher, see if that works

	if err != nil {
		t.Fatal(err)
	}

	m := make([]byte, cipher.BlockSize())
	c := make([]byte, cipher.BlockSize())

	cipher.Encrypt(c, m)

	if len(m) != len(c) {
		t.Error("lengths don't match")
	}

	// does this test even make sense? what's the probability that it fails?
	// how can it be corrected so that it makes sense? (false negative should be as improbable as guessing the key)
	for i, _ := range m {
		if m[i] == c[i] {
			t.Error("some of the encrypted characters equal")
		}
	}

	d := make([]byte, cipher.BlockSize())
	cipher.Decrypt(d, c)

	if len(d) != len(m) {
		t.Error("lengths don't match")
	}

	for i, _ := range m {
		if d[i] != m[i] {
			t.Error(fmt.Sprintf("D(E(m[%d]))=%d != m[%d]=%d", i, d[i], i, m[i]))
		}
	}
}
