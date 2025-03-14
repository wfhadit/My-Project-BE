package services

import "my-project-be/features/product"

type service struct {
	model product.ProductModel
}

func ProductService(pm product.ProductModel) product.ProductService {
	return &service{model: pm}
}

func (s *service) CreateProduct(newData product.Product) ( product.Product, error) {
	result,err := s.model.CreateProduct(newData)
	if err != nil {
		return newData, err
	}
	return result, nil
}

func (s *service) GetAllProducts() ([]product.Product, error) {
	result, err := s.model.GetAllProducts()
	if err != nil {
		return []product.Product{}, err
	}
	return result, nil
}

func (s *service) GetProductByID(productID uint) (product.Product, error) {
	result, err := s.model.GetProductByID(productID)
	if err != nil {
		return product.Product{}, err
	}
	return result, nil
}