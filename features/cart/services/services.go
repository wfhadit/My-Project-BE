package services

import (
	"my-project-be/features/cart"
	"my-project-be/middlewares"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model cart.CartModel
}

func CartService(cm cart.CartModel) cart.CartService {
	return &service{model: cm}
}

func (s *service) AddCart(token *jwt.Token, newData cart.Cart)  error {
	userid,_,_ := middlewares.DecodeToken(token)
	if userid == 0 {
		return  nil
	}
	err := s.model.AddCart(userid, newData)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetCart(token *jwt.Token) ([]cart.Cart, error) {
	userid,_,_ := middlewares.DecodeToken(token)
	if userid == 0 {
		return []cart.Cart{}, nil
	}
	result,err := s.model.GetCart(userid)
	if err != nil {
		return []cart.Cart{}, err
	}
	return result, nil
}

func (s *service) DeleteCart(token *jwt.Token) error {
	userid,_,_ := middlewares.DecodeToken(token)
	if userid == 0 {
		return  nil
	}
	err := s.model.DeleteCart(userid)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteCartByID(token *jwt.Token, productID uint) error {
	userid,_,_ := middlewares.DecodeToken(token)
	if userid == 0 {
		return  nil
	}
	err := s.model.DeleteCartByID(userid, productID)
	if err != nil {
		return err
	}
	return nil
}