package user

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
}

type UserModel interface {
	Register(newData User) error
	Login(email string) (User, error)
}
