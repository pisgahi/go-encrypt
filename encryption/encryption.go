package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func Encrypt() {
	plainText, err := os.ReadFile("assets/plaintext.txt")
	if err != nil {
		log.Fatal("Error reading file", err.Error())
	}

	key, err := os.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatal("Error reading file", err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce err: %v", err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	err = os.WriteFile("assets/outputs/encrypted/ciphertext.bin", cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
