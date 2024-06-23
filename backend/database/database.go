package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		fmt.Println("Databse URL is empty")
	}

	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(DB_URL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connection to DB", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging DB")
	}

	fmt.Println("Connected to DB")
	return client
}
