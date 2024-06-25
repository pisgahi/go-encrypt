package models

type UserSecret struct {
	PlainText string `json:"plainText"`
	Key       string `json:"key"`
}
type UserKey struct {
	Key string `json:"key"`
}

type DecryptedSecret struct {
	DecipheredText string `json:"decipheredText"`
}
