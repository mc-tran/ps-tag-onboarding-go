package services

import (
	"context"
	"fmt"
	"log"

	"github.com/mc-tran/ps-tag-onboarding-go/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	mongoclient *mongo.Client
	context     context.Context
}

func NewUserService() *UserService {

	ctx := context.TODO()
	client := CreateMongoClient(ctx)

	return &UserService{mongoclient: client, context: ctx}
}

func CreateMongoClient(c context.Context) *mongo.Client {

	//db is docker-compose service name
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@db:27017")
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

func (us *UserService) AddUser(p *data.User) {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	insertedDocument := bson.M{
		"id":        p.ID,
		"firstname": p.FirstName,
		"lastname":  p.LastName,
		"email":     p.Email,
		"age":       p.Age,
	}

	_, err := collection.InsertOne(us.context, insertedDocument)

	log.Printf("inserted")

	if err != nil {
		panic(err)
	}
}

func (us *UserService) GetUser(id string) data.User {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	var user data.User
	err := collection.FindOne(us.context, bson.M{"id": id}).Decode(&user)

	if err != nil {
		panic(err)
	}

	fmt.Println("get user")

	return user
}

func (us *UserService) DoesUserExist(firstname string, lastname string) bool {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	filter := bson.M{"firstname": firstname, "lastname": lastname}

	count, err := collection.CountDocuments(us.context, filter)

	if err != nil {
		panic(err)
	}

	return count > 0
}
