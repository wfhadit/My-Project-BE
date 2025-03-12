package data

type User struct {
	ID           uint `gorm:"primary_key,auto_increment"`
	Nama         string
	Email        string
	Password     string
	JenisKelamin string
	TanggalLahir string
	NomorHP      string
	Alamat       string
	Foto         string
}