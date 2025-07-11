package types

import (
	"time"
)

type User struct{
	ID int `json"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"` // URL to the image
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
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

type CreateProductPayLoad struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"` // URL to the image
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) (error)
}

type ProductStore interface{
	GetProductByID(id int) (*Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(product *Product) error
}