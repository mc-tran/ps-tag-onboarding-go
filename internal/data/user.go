package data

import (
	"encoding/json"
	"io"
	"net/mail"

	"github.com/mc-tran/ps-tag-onboarding-go/internal/constants"
)

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Age       int    `json:"age" bson:"age"`
}

type Users []*User

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

	_, err := mail.ParseAddress(email)
	if err != nil {
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
