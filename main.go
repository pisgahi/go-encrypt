package main

import (
	decrypt "github.com/pisgahi/go-encrypt/decryption"
	encrypt "github.com/pisgahi/go-encrypt/encryption"
)

func main() {
	encrypt.Encrypt()
	decrypt.Decrypt()
}
