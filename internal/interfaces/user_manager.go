package interfaces

import (
	"github.com/mc-tran/ps-tag-onboarding-go/internal/data"
)

type UserManager interface {
	GetUser(id string) data.User
	AddUser(p *data.User) string
	DoesUserExist(firstname string, lastname string) bool
}
