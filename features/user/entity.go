package user

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID           uint
	Nama         string
	Email        string
	Password     string
	JenisKelamin string
	NomorHP      string
	Alamat       string
	Foto         string
}
type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	KeepLogin(token *jwt.Token) (User, string, error)
}

type UserModel interface {
	Register(newData User) error
	Login(email string) (User, error)
	GetUserByID(id uint) (User, error)
}
