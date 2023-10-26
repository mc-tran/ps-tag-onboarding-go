package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/constants"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/data"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/interfaces"
)

type UsersHandler struct {
	l           *log.Logger
	userService interfaces.UserRepository
}

func NewUsersHandler(l *log.Logger, userService interfaces.UserRepository) *UsersHandler {
	return &UsersHandler{l, userService}
}

func (p *UsersHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {

}

func (p *UsersHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Users")

	vars := mux.Vars(r)
	id, _ := vars["id"]

	user := p.userService.GetUser(id)

	user.ToJSON(rw)
}

func (p *UsersHandler) AddUsers(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Users")

	user := r.Context().Value(KeyUser{}).(data.User)

	p.userService.AddUser(&user)
}

type KeyUser struct{}

func (p UsersHandler) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.User{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Error reading user", http.StatusBadRequest)
			return
		}

		validationErr := prod.ValidateFields()

		if p.userService.DoesUserExist(prod.FirstName, prod.LastName) {
			validationErr = append(validationErr, constants.Error_Name_Unique)
		}

		if len(validationErr) > 0 {

			var errors = strings.Join(validationErr[:], ",")

			p.l.Println("[ERROR] validating user", errors)

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
