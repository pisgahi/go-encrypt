package main

import (
	decrypt "github.com/pisgahi/go-encrypt/decryption"
)

func main() {
	decrypt.Decrypt("this_is_32_byte_key_for_AES_256!")
}
