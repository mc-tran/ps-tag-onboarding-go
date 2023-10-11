package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/minh/constants"
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

	vars := mux.Vars(r)
	id, _ := vars["id"]

	userservice := services.NewUserService()
	user := userservice.GetUser(id)

	user.ToJSON(rw)
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

		validationErr := prod.ValidateFields()

		userservice := services.NewUserService()

		if userservice.DoesUserExist(prod.FirstName, prod.LastName) {
			validationErr = append(validationErr, constants.Error_Name_Unique)
		}

		if len(validationErr) > 0 {
			p.l.Println("[ERROR] validating user", err)

			var errors = strings.Join(validationErr[:], ",")
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", errors),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUser{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
