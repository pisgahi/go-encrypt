package main

import (
	databse "github.com/pisgahi/go-encrypt/database"
	encrypt "github.com/pisgahi/go-encrypt/encryption"
)

func main() {
	databse.GetClient()
	encrypt.Encrypt("This should work", "this_is_32_byte_key_for_AES_256!")
}
