package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minh/data"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	mongoclient *mongo.Client
}

func NewUserService() *UserService {
	client := CreateMongoClient()

	return &UserService{mongoclient: client}
}

func CreateMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017/?authSource=admin")
	mongoClient, err := mongo.Connect(ctx, clientOptions)

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("connection error :%v", err)
		panic(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		panic(err)
	}

	fmt.Println("ping success")

	return mongoClient
}

func (us *UserService) AddUser(p *data.User) {

	fmt.Println("adding user")
}

func (us *UserService) GetUser(p *data.User) {

	fmt.Println("adding user")
}
