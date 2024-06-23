package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"

	"github.com/pisgahi/go-encrypt/database"
)

func Decrypt(key string) string {
	keyBytes := []byte(key)

	cipherText := database.GetSecret(key)

	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		log.Fatalf("base64 decode err: %v", err.Error())
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := cipherTextBytes[:gcm.NonceSize()]
	cipherTextBytes = cipherTextBytes[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	decrypted := string(plainText)

	return decrypted
}
