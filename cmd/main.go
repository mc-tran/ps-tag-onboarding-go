package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/handlers"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/repository"
)

func main() {

	l := log.New(os.Stdout, "test-api", log.LstdFlags)

	mongoUri := os.Getenv("APP_MONGO_CONNECTION_STRING")

	mongoClient := repository.CreateMongoClient(context.TODO(), mongoUri)

	userRepository := repository.NewUserRepository(mongoClient)

	uh := handlers.NewUsersHandler(l, userRepository)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/users", uh.GetUsers)
	getRouter.HandleFunc("/find/{id}", uh.GetUser)
	getRouter.Use(handlers.ErrorHandler)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/save", uh.AddUsers)
	postRouter.Use(uh.MiddlewareValidateUser)
	postRouter.Use(handlers.ErrorHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
