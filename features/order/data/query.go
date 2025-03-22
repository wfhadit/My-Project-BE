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

func (m *model) CreateOrder(newOrder order.Order) (order.Order, error) {
	errOrder := m.connection.Create(&newOrder).Error
	if errOrder != nil {
		return order.Order{}, errOrder
	}
	return newOrder, nil
}

func (m *model) GetOrderByUniqueID(uniqueID string, userid uint, newStatus string) (order.Order, error) {
	result := order.Order{}
	errUpdate := m.connection.Model(&order.Order{}).Where("order_unique_id = ?", uniqueID).UpdateColumn("status", newStatus).Error
	if errUpdate != nil {
		return order.Order{}, errUpdate
	}
	errOrder := m.connection.Preload("Items").Where("order_unique_id = ?", uniqueID).First(&result).Error
	if errOrder != nil {
		return order.Order{}, errOrder
	}
	return result, nil
}

func (m *model) GetAllOrders(userid uint) ([]order.Order, error) {
	result := []order.Order{}
	errOrder := m.connection.Preload("Items").Where("user_id = ?", userid).Find(&result).Error
	if errOrder != nil {
		return []order.Order{}, errOrder
	}
	return result, nil
}

func (m *model) GetLastOrder(userid uint) (order.Order, error) {
	result := order.Order{}
	errOrder := m.connection.Preload("Items").Where("user_id = ?" , userid).Last(&result).Error
	if errOrder != nil {
		return order.Order{}, errOrder
	}
	return result, nil
}