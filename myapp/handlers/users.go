package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/minh/data"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (p *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Users")

	listUsers := data.GetUsers()

	err := listUsers.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert products to JSON", http.StatusInternalServerError)
	}
}

func (p *Users) GetUser(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Users")

	vars := mux.Vars(r)
	id, _ := vars["id"]

	user, _ := data.GetUserByID(id)

	user.ToJSON(rw)
}

func (p *Users) AddUsers(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Users")

	user := &data.User{}

	err := user.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to decode JSON derr", http.StatusBadRequest)
	}

	data.AddUser(user)
}

type KeyUser struct{}

func (p Users) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.User{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyUser{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
