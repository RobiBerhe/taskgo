package user

import "time"

type User struct {
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
