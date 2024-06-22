package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func Decrypt(cipherText []byte, key string) {
	keyBytes := []byte(key)
	cipherTextBytes := []byte(cipherText)

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

	err = os.WriteFile("assets/outputs/decrypted/plaintext.txt", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}
