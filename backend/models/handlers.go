package models

type UserSecret struct {
	CipherText string
	Key        string
}
type UserKey struct {
	Key string
}

type DecryptedSecret struct {
	DecipheredText string `json:"decipheredText"`
}
