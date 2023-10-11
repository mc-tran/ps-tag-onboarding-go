package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/minh/data"
	"github.com/minh/services"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (p *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {

}

func (p *Users) GetUser(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Users")

	// vars := mux.Vars(r)
	// id, _ := vars["id"]

	// user, _ := data.GetUserByID(id)
	// user.ToJSON(rw)
}

func (p *Users) AddUsers(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Users")

	user := r.Context().Value(KeyUser{}).(data.User)

	userservice := services.NewUserService()
	userservice.AddUser(&user)
}

type KeyUser struct{}

func (p Users) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.User{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Error reading user", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating user", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUser{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
