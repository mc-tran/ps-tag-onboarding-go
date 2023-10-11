package data

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator"

	"github.com/minh/constants"
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

func (p *User) ValidateFields() []string {

	var e []string

	if p.FirstName == "" {
		e = append(e, constants.Error_Name_Required)
	}

	if p.LastName == "" {
		e = append(e, constants.Error_Name_Required)
	}

	if p.Email == "" {
		e = append(e, constants.Error_Email_Required)
	} else if !validateEmail(p.Email) {
		e = append(e, constants.Error_Email_Format)
	}

	if p.Age < 18 {
		e = append(e, constants.Error_Age_Minimum)
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
