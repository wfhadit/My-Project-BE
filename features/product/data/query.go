package data

import (
	"my-project-be/features/product"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func ProductModel(db *gorm.DB) product.ProductModel {
	return &model{connection: db}
}

func (m *model) CreateProduct(newData product.Product) ( product.Product, error) {
	err := m.connection.Create(&newData).Error
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *model) GetAllProducts() ([]product.Product, error) {
	result := []product.Product{}
	err := m.connection.Find(&result).Error
	if err != nil {
		return []product.Product{}, err
	}
	return result, nil
}

func (m *model) GetProductByID(productID uint) (product.Product, error) {
	result := product.Product{}
	err := m.connection.Where("id = ?", productID).First(&result).Error
	if err != nil {
		return product.Product{}, err
	}
	return result, nil
}