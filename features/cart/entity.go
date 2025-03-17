package cart

import "github.com/golang-jwt/jwt/v5"

type Cart struct {
	ProductID    uint
	ProductNama  string
	ProductImage string
	ProductPrice uint
	Quantity     uint
}

type CartService interface {
	AddCart(token *jwt.Token,newData Cart) error
	GetCart(token *jwt.Token)([]Cart,error)
	DeleteCartByID(token *jwt.Token, productID uint) error
	DeleteCart(token *jwt.Token) error
}

type CartModel interface {
	AddCart(userid uint, newData Cart) error
	GetCart(userid uint)([]Cart,error)
	DeleteCartByID(userid uint, productID uint) error
	DeleteCart(userid uint) error
}