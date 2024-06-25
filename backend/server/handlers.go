package server

import (
	"encoding/json"
	"log"
	"net/http"

	decrypt "github.com/pisgahi/go-encrypt/decryption"
	encrypt "github.com/pisgahi/go-encrypt/encryption"
	"github.com/pisgahi/go-encrypt/models"
)

func serverGreeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello go-encrpt"))
}

func sendSecretHandler(w http.ResponseWriter, r *http.Request) {
	var newUserSecret models.UserSecret
	err := json.NewDecoder(r.Body).Decode(&newUserSecret)
	if err != nil {
		log.Fatal("Error decoding secret", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	encrypt.Encrypt(newUserSecret.PlainText, newUserSecret.Key)
	w.Write([]byte("Secret added"))
	w.WriteHeader(http.StatusOK)
}

func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	var userKey models.UserKey
	err := json.NewDecoder(r.Body).Decode(&userKey)
	if err != nil {
		log.Fatal("Erorr decoding key", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	decipheredStr := decrypt.Decrypt(userKey.Key)
	decryptedSecret := models.DecryptedSecret{
		DecipheredText: decipheredStr,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(decryptedSecret)
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
