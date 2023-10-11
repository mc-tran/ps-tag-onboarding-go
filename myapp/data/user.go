package data

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

type User struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"gt=18"`
}

type Users []*User

func (p *User) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("email", validateEmail)
	return validate.Struct(p)
}

func validateEmail(fl validator.FieldLevel) bool {

	re := regexp.MustCompile(`.+\@.+\..+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
