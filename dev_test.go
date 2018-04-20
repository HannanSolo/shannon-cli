package main_test

import (
	"fakecypher"
	"fmt"
	"testing"
)

// this test should really have been broken down into several tests
// each test can be used to check your assumptions about how things work
// when writing tests for new code, you can use tests to assert your assumptions about how it *should* work
func TestBlockInterface(t *testing.T) {
	key := make([]byte, 16)
	// _, err := rand.Read(key)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	cipher, err := fakecypher.NewCipher(key) // TODO replace this with FakeCipher, see if that works

	if err != nil {
		t.Fatal(err)
	}

	m := make([]byte, cipher.BlockSize())
	m[0] = 1
	// _, err = rand.Read(m)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	c := make([]byte, cipher.BlockSize())

	cipher.Encrypt(c, m)

	if len(m) != len(c) {
		t.Error("lengths should match")
	}

	// does this test even make sense? what's the probability that it fails?
	// how can it be corrected so that it makes sense? (false negative should be as improbable as guessing the key)
	if string(m) == string(c) {
		t.Error("failed encryption")
	}

	fmt.Println(m)

	fmt.Println(c)

	d := make([]byte, cipher.BlockSize())
	cipher.Decrypt(d, c)

	if len(d) != len(m) {
		t.Error("lengths should match")
	}

	fmt.Printf("%v\n", d)
	for i, _ := range m {
		if d[i] != m[i] {
			t.Error(fmt.Sprintf("D(E(m[%d]))=%d != m[%d]=%d", i, d[i], i, m[i]))
		}
	}
}
