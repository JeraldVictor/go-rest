package types

import "time"

type Userstore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type RegisterUserPayload struct {
	Username string `json:"userName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=120"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"userName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
