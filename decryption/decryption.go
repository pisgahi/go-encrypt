package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func Decrypt() {
	cipherText, err := os.ReadFile("assets/outputs/encrypted/ciphertext.bin")
	if err != nil {
		log.Fatal(err)
	}

	key, err := os.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	err = os.WriteFile("assets/outputs/decrypted/plaintext.txt", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
