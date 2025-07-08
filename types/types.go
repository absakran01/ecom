package types

import "time"

type User struct{
	ID int `json"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayLoad struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) (error)
}