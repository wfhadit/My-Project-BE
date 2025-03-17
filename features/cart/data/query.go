package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"my-project-be/features/cart"

	"github.com/redis/go-redis/v9"
)

type model struct {
	connection *redis.Client     
}

func CartModel(rds *redis.Client) cart.CartModel {
	return &model{connection: rds}
}

func (m *model) AddCart(userid uint, newData cart.Cart)  error {
	rdb:= m.connection
	dataJSON, err := json.Marshal(newData)
	if err != nil {
		return err
	}
	rdb.HSet(context.Background(), fmt.Sprintf("cart:%d", userid), fmt.Sprintf("%d", newData.ProductID), string(dataJSON))
	return nil
}

func (m *model) GetCart(userid uint) ([]cart.Cart, error) {
	rdb := m.connection
	result, err := rdb.HGetAll(context.Background(), fmt.Sprintf("cart:%d", userid)).Result()
	if err != nil {
		return []cart.Cart{}, err
	}
	var cartItems []cart.Cart
	for _, item := range result {
		var cartItem cart.Cart
		err := json.Unmarshal([]byte(item), &cartItem)
		if err != nil {
			log.Println("Error unmarshaling:", err)
			return []cart.Cart{}, err
		}
		
		cartItems = append(cartItems, cartItem)
	}


	return cartItems, nil
}

func (m *model) DeleteCartByID(userid uint, productID uint) error {
	rdb := m.connection
	err := rdb.HDel(context.Background(), fmt.Sprintf("cart:%d", userid), fmt.Sprintf("%d", productID)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m *model) DeleteCart(userid uint) error {
	rdb := m.connection
	err := rdb.Del(context.Background(), fmt.Sprintf("cart:%d", userid)).Err()
	if err != nil {
		return err
	}
	return nil
}
