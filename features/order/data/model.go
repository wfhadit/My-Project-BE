package data

type Order struct {
	ID            uint `gorm:"primary_key,auto_increment"`
	UserID        uint
	TotalPrice    uint
	Status        string
	PaymentMethod string
	VANumber      string
	OrderUniqueID string
	CreatedAt     string
	Items         []OrderItem `gorm:"foreign_key:OrderID"`
}

type OrderItem struct {
	ID           uint `gorm:"primary_key,auto_increment"`
	OrderID      uint
	ProductID    uint
	ProductName  string
	ProductImage string
	ProductPrice uint
	Quantity     uint
}