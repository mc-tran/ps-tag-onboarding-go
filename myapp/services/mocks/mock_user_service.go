package mockservice

import (
	"fmt"

	"github.com/minh/data"
)

type MockUserService struct{}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (us *MockUserService) GetUserByID(id string) (*data.User, error) {
	for _, user := range userList {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %s not found", id)
}

func (us *MockUserService) AddUser(p *data.User) {
	userList = append(userList, p)
}

func (us *MockUserService) GetUsers() data.Users {
	return userList
}

var userList = []*data.User{
	&data.User{
		ID:        "111",
		FirstName: "Michael",
		LastName:  "Scott",
		Email:     "aaa@aaa.com",
		Age:       23,
	},
	&data.User{
		ID:        "222",
		FirstName: "Tom",
		LastName:  "Van",
		Email:     "bbb@bbb.com",
		Age:       20,
	},
}
