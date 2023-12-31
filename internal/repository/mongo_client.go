package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(c context.Context, connectionString string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(connectionString)
	mongoClient, err := mongo.Connect(c, clientOptions)

	if err != nil {
		log.Fatalf("connection error :%v", err)
		panic(err)
	}

	err = mongoClient.Ping(c, nil)
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		panic(err)
	}

	fmt.Println("ping success")

	return mongoClient
}
