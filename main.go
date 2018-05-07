package main

import (
	"encrypt"
	"fmt"
)

/*


Example usage:

Encrypt:

$ foo -p secret.txt
Password: 01234decafbaddeadbeef831751959

AES
(Anubis)
(BlowFish)
Cast5
Cast6
Serpent
Tnepres
TwoFish

concerns with interactive use:

- command line argument leak to the unix namespace (e.g. visible to top/ps, and any other tool, on linux via /proc filesystem)
- interactive password input requires putting tty in non echo mode (see go packages for securely reading password input, also requires double entry for verification)
- user determined passwords are likely to be low entropy, and therefore easy to brute force, a key stretching function with a unique nonce (salt)
- example real world setup, linux unified key setup - key concepts are the master key, key slots, key derivation from passwords, and anti foresnsics features (not relevant for your application)
  http://tomb.dyne.org/Luks_on_disk_format.pdf


- conservative cipher suite options you can use:
  - symmetric ciphers - aes256, chacha, serpent
  - hash function - sha256, sha512, sha3, blake2
  - key derivation - argon2, scrypt, pdfk2
  - symmetric cipher mode - look up what *NOT* to do (e.g. ECB mode), I would reccomend an authenticated encryption and data (AEAD) mode, such as GCM or EAX -- these modes include intergity protection in the encryption
    (modes that do not have integrity protection, CBC, CTR, ... )

therefore for initial version we should just generate a random key



Decrypt:

$ foo -d secret.txt.enc
Password: 01234decafbaddeadbeef831751959

*/

//var config config.Config // package config       type Config struct { Algorithm string }

// var decrypt bool

// func init() {
// 	flag.BoolVar(&decrypt, "d", false, "Decryption mode")
// }

func main() {
	//Caesars CLI, Lets do this
	// flag.Parse()
	// args := flag.Args()

	// if len(args) < 1 {
	// 	fmt.Printf("Too few arguments \n")
	// 	fmt.Printf("Usage of %s: \n", os.Args[0])
	// 	flag.PrintDefaults()
	// 	return
	// }
	cipherText := encrypt.EncryptAES([]byte("hello, world"), "password")

	fmt.Println(string(cipherText))

	plainText := encrypt.DecryptAES(cipherText, "password")
	fmt.Println(string(plainText))

	/// at this point we assume all options are valid

	//file := args[0]

	// if !decrypt {
	// 	encryptor.EncryptFile(config, file)
	// } else {
	// 	decryptor.DecryptFile(config, file)
	// }
}
