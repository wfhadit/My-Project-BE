package data

type Product struct {
	ID          uint `gorm:"primary_key,auto_increment"`
	Nama        string
	Brand       string
	Category    string
	Price       uint
	Amount      uint
	Description string
	Image       string
}