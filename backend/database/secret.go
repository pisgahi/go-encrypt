package database

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	"github.com/pisgahi/go-encrypt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSecret(cipherText []byte, key string) error {
	client := GetClient()

	cipherTextBase64 := base64.StdEncoding.EncodeToString(cipherText)

	newSecret := models.Secret{
		Id:         primitive.NewObjectID().Hex(),
		CipherText: cipherTextBase64,
		Key:        key,
	}

	collection := client.Database("go-encrypt").Collection("secrets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newSecret)
	if err != nil {
		log.Println("Error add comment to DB", err)
	}

	return nil
}

func GetSecret(key string) string {
	client := GetClient()

	collection := client.Database("go-encrypt").Collection("secrets")

	var result models.Secret
	filter := bson.M{"key": key}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}

	if result.Key != key {
		log.Fatalf("Provided key does not match the key used for encryption")
	}

	return result.CipherText
}
