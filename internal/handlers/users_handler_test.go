package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/data"
	"github.com/mc-tran/ps-tag-onboarding-go/internal/services"
	"github.com/stretchr/testify/assert"
)

func UsersHandler_FindExistingUser(t *testing.T) {

	logger := log.New(bytes.NewBufferString(""), "", 0)

	mockservice := services.NewMockUserService()

	user := &data.User{
		ID:        "111",
		FirstName: "Michael",
		LastName:  "Scott",
		Email:     "aaa@aaa.com",
		Age:       23,
	}

	// Create a new router
	router := mux.NewRouter()

	// Create a new user handler
	userHandler := NewUsersHandler(logger, mockservice)

	// Register the user handler routes
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	// Create a new request
	req, err := http.NewRequest("GET", "/users/111", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "111",
	}

	req = mux.SetURLVars(req, vars)

	// Call the GetUser function
	userHandler.GetUser(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseUser data.User
	err = json.Unmarshal(rr.Body.Bytes(), &responseUser)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user, &responseUser)
}
