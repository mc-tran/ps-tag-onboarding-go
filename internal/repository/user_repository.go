package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mc-tran/ps-tag-onboarding-go/internal/constants"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/customerrors"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoclient *mongo.Client
	context     context.Context
}

func NewUserRepository(client *mongo.Client) *UserRepository {

	ctx := context.TODO()

	return &UserRepository{mongoclient: client, context: ctx}
}

func (us *UserRepository) AddUser(p *data.User) string {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	insertedDocument := bson.M{
		"firstname": p.FirstName,
		"lastname":  p.LastName,
		"email":     p.Email,
		"age":       p.Age,
	}

	i, err := collection.InsertOne(us.context, insertedDocument)

	log.Printf("inserted")

	if err != nil {
		panic(err)
	}

	inserted := i.InsertedID.(primitive.ObjectID).String()

	return inserted
}

func (us *UserRepository) GetUser(id string) (data.User, error) {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	objID, _ := primitive.ObjectIDFromHex(id)

	var user data.User
	err := collection.FindOne(us.context, bson.M{"_id": objID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return data.User{}, &customerrors.UserNotFoundError{Message: constants.Error_User_Not_Found}
		}

		panic(err)
	}

	fmt.Println("get user")

	return user, nil
}

func (us *UserRepository) DoesUserExist(firstname string, lastname string) bool {

	database := us.mongoclient.Database("user")
	collection := database.Collection("userdetails")

	filter := bson.M{"firstname": firstname, "lastname": lastname}

	var user data.User
	err := collection.FindOne(us.context, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		panic(err)
	}

	return true
}
