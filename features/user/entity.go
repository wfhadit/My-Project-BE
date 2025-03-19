package user

import (
	"mime/multipart"
	"my-project-be/features/cart"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           uint
	Nama         string
	Email        string
	Password     string
	JenisKelamin string
	TanggalLahir string
	NomorHP      string
	Alamat       string
	Foto         string
}



type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string,[]cart.Cart, error)
	KeepLogin(token *jwt.Token) (User, string,[]cart.Cart, error)
	Update(token *jwt.Token, newData User, file *multipart.FileHeader ) (User, error)
}

type UserModel interface {
	Register(newData User) error
	Login(email string) (User, error)
	GetUserByID(id uint) (User, error)
	Update(id uint, newData User) (User, error)
}
