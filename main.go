package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongo_auth := os.Getenv("MONGODB_AUTH")
	mongo_username := os.Getenv("MONGODB_USER")
	mongo_password := os.Getenv("MONGODB_PASSWORD")
	clientOptions := options.Client()
	if mongo_auth == "true" {
		credential := options.Credential{
			Username: mongo_username,
			Password: mongo_password,
		}
		clientOptions = options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT")).SetAuth(credential)
	} else {
		clientOptions = options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT"))
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb respond OK ")
}
