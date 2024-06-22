package database

import (
	"context"
	"fmt"
	"time"
)

type Secret struct {
	id string
}

func AddSecret() error {
	client := GetClient()

	newSecret := Secret{
		id: "test",
	}

	collection := client.Database("go-encrypt").Collection("secrets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newSecret)
	if err != nil {
		fmt.Println("Error add comment to DB", err)
	}
	return nil

}
