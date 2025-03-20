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

func (m *model) GetAllProducts(offset int, category, brand, sort, q string) ([]product.Product, int64,error) {
	result := []product.Product{}
	var total int64
	query := m.connection.Model(&product.Product{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if brand != "" {
		query = query.Where("brand = ?", brand)
	}
	if q != "" {
		query = query.Where("nama LIKE ? OR brand LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	query.Count(&total)
	query = query.Limit(10).Offset(offset)
	if sort == "asc" {
		query = query.Order("price ASC")
	} else if sort == "desc" {
		query = query.Order("price DESC")
	}
	err := query.Find(&result).Error
	if err != nil {
		return []product.Product{}, 0, err
	}
	return result, total,nil
}

func (m *model) GetProductByID(productID uint) (product.Product, error) {
	result := product.Product{}
	err := m.connection.Where("id = ?", productID).First(&result).Error
	if err != nil {
		return product.Product{}, err
	}
	return result, nil
}

