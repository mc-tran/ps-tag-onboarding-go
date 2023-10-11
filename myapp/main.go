package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/minh/handlers"
)

func main() {

	l := log.New(os.Stdout, "minh-api", log.LstdFlags)

	uh := handlers.NewUsers(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/users", uh.GetUsers)
	getRouter.HandleFunc("/find/{id:[0-9]+}", uh.GetUser)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/save", uh.AddUsers)
	postRouter.Use(uh.MiddlewareValidateUser)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
