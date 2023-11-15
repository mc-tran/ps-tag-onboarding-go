package services

import (
	"github.com/mc-tran/ps-tag-onboarding-go/internal/data"
)

type MockUserService struct{}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (us *MockUserService) GetUser(id string) data.User {
	// for _, user := range userList {
	// 	if user.ID == id {
	// 		return *user
	// 	}
	// }
	return data.User{}
}

func (us *MockUserService) AddUser(p *data.User) {
	userList = append(userList, p)
}

func (us *MockUserService) DoesUserExist(firstname string, lastname string) bool {
	for _, user := range userList {
		if user.FirstName == firstname && user.LastName == lastname {
			return true
		}
	}
	return false
}

func (us *MockUserService) GetUsers() data.Users {
	return userList
}

var userList = []*data.User{
	&data.User{
		FirstName: "Michael",
		LastName:  "Scott",
		Email:     "aaa@aaa.com",
		Age:       23,
	},
	&data.User{
		FirstName: "Tom",
		LastName:  "Van",
		Email:     "bbb@bbb.com",
		Age:       20,
	},
}
