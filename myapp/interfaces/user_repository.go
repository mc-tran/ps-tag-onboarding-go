package interfaces

import (
	"github.com/mc-tran/ps-tag-onboarding-go/data"
)

type UserRepository interface {
	GetUser(id string) data.User
	AddUser(p *data.User)
	DoesUserExist(firstname string, lastname string) bool
}
