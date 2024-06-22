package main

import (
	encrypt "github.com/pisgahi/go-encrypt/encryption"
)

func main() {
	encrypt.Encrypt("This should work", "this_is_32_byte_key_for_AES_256!")
}
