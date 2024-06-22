package main

import (
	decrypt "github.com/pisgahi/go-encrypt/decryption"
	encrypt "github.com/pisgahi/go-encrypt/encryption"
)

func main() {
	encrypt.Encrypt("New one", "this_is_32_byte_key_for_AES_256!")
	decrypt.Decrypt("this_is_32_byte_key_for_AES_256!")
}
