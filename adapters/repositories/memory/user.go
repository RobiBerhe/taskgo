package memory

import (
	"errors"
	"log"
	"taskgo/internal/core/user"
	"taskgo/pkg/exception"

	"github.com/google/uuid"
)

type User struct {
}

// Create implements repositories.User.
func (u *User) Create(user *user.User) *exception.Exception {
	log.Printf("the user has been created : %v", user)
	return exception.New(400, errors.New("something went wrong"), "sorry, something went wrong", "something went wrong")
}

// GetById implements repositories.User.
func (u *User) GetById(userId uuid.UUID) {
	panic("unimplemented")
}

func NewUserRepo() user.UserRepository {
	return &User{}
}
