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
	Email     string `json:"email" validate:"required"`
	Age       int    `json:"age" validate:"gt=18"`
}

type Users []*User

func (p *User) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}

func (p *User) ValidateNew() []string {

	var e []string

	if p.FirstName == "" {
		e = append(e, "First Name is required")
	}

	if p.LastName == "" {
		e = append(e, "Last Name is required")
	}

	if p.Email == "" {
		e = append(e, "Last Name is required")
	} else if !validateEmail(p.Email) {
		e = append(e, "Email is invalid")
	}

	if p.Age < 18 {
		e = append(e, "Age must be greater than 18")
	}

	return e
}

func validateEmail(email string) bool {

	re := regexp.MustCompile(`.+\@.+\..+`)
	matches := re.FindAllString(email, -1)

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
