package data

import (
	"testing"

	"github.com/mc-tran/ps-tag-onboarding-go/internal/constants"

	"github.com/stretchr/testify/assert"
)

func TestUserFieldsValidation_ValidUser(t *testing.T) {

	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       25,
	}

	// Assert
	errors := user.ValidateFields()
	assert.Empty(t, errors)

}

func TestUserFieldsValidation_InvalidEmail(t *testing.T) {

	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.example.com",
		Age:       25,
	}

	// Assert
	errors := user.ValidateFields()
	assert.Equal(t, errors[0], constants.Error_Email_Format)
}

func TestUserFieldsValidation_EmptyEmail(t *testing.T) {

	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	// Assert
	errors := user.ValidateFields()
	assert.Equal(t, errors[0], constants.Error_Email_Required)
}

func TestUserFieldsValidation_InvalidAge(t *testing.T) {

	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       17,
	}

	// Assert
	errors := user.ValidateFields()
	assert.Equal(t, errors[0], constants.Error_Age_Minimum)
}

func TestUserFieldsValidation_MissingName(t *testing.T) {

	user := &User{
		Email: "john.doe@example.com",
		Age:   17,
	}

	// Assert
	errors := user.ValidateFields()
	assert.Equal(t, errors[0], constants.Error_Name_Required)
}
