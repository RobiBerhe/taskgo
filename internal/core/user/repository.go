package user

import (
	"taskgo/pkg/exception"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) *exception.Exception
	GetById(userId uuid.UUID)
}
