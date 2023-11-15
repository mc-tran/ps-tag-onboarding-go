package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/handlers"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/services"
)

func main() {

	l := log.New(os.Stdout, "minh-api", log.LstdFlags)

	userService := services.NewUserService()
	uh := handlers.NewUsersHandler(l, userService)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/users", uh.GetUsers)
	getRouter.HandleFunc("/find/{id:[0-9]+}", uh.GetUser)
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
