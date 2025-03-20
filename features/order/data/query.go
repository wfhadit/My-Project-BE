package data

import (
	"my-project-be/features/order"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func OrderModel(db *gorm.DB) order.OrderModel {
	return &model{connection: db}
}

func (m *model) CreateOrder(newOrder order.Order, userid uint) (order.Order, error) {
	errOrder := m.connection.Create(&newOrder).Error
	if errOrder != nil {
		return order.Order{}, errOrder
	}
	return newOrder, nil
}