package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"github.com/pisgahi/go-encrypt/database"
)

func Encrypt(plainText string, key string) {
	keyBytes := []byte(key)
	plainTextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
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

	cipherText := gcm.Seal(nonce, nonce, plainTextBytes, nil)

	database.AddSecret(cipherText, key)
}
