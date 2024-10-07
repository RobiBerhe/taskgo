package user

import (
	"time"
)

type UserDTO struct {
	Name      string `json:"name" validate:"required,min=3,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4,max=30"`
	CreatedAt time.Time
}

type CreateUserRequest struct {
	User UserDTO `json:"user" validate:"required"`
}

func (cur *CreateUserRequest) Validate() []error {
	return nil
}

func (cur *CreateUserRequest) ToUser() *User {
	return &User{
		Name:     cur.User.Name,
		Email:    cur.User.Email,
		Password: cur.User.Password,
	}
}
