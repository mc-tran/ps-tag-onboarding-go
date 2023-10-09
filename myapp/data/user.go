package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required"`
	Age       int    `json:"age" validate:"gt=18"`
}

type Users []*User

func GetUsers() Users {
	return userList
}

func (p *User) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}

func GetUserByID(id string) (*User, error) {
	for _, user := range userList {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %s not found", id)
}

func (p *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Users) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddUser(p *User) {
	userList = append(userList, p)
}

var userList = []*User{
	&User{
		ID:        "111",
		FirstName: "Erling",
		LastName:  "Haaland",
		Email:     "aaa@aaa.com",
		Age:       23,
	},
	&User{
		ID:        "222",
		FirstName: "Micky",
		LastName:  "Van De Ven",
		Email:     "bbb@bbb.com",
		Age:       20,
	},
}
