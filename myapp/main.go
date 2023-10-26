package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/minh/handlers"
	"github.com/minh/services"
)

func main() {

	l := log.New(os.Stdout, "minh-api", log.LstdFlags)

	userService := services.NewUserService()
	uh := handlers.NewUsersHandler(l, userService)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/users", uh.GetUsers)
	getRouter.HandleFunc("/find/{id:[0-9]+}", uh.GetUser)
	getRouter.Use(ErrorHandler)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/save", uh.AddUsers)
	postRouter.Use(uh.MiddlewareValidateUser)
	postRouter.Use(ErrorHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}

func ErrorHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				http.Error(w, fmt.Sprintf("An error occured: %s", err), http.StatusInternalServerError)

			}
		}()

		next.ServeHTTP(w, r)

	})

}
