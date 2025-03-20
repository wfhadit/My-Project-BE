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

func (s *service) GetAllProducts(offset int, category, brand, sort, q string) ([]product.Product, int64,error) {
	result, total, err := s.model.GetAllProducts(offset, category, brand,  sort, q)
	if err != nil {
		return []product.Product{}, 0,err
	}
	return result, total, nil
}

func (s *service) GetProductByID(productID uint) (product.Product, error) {
	result, err := s.model.GetProductByID(productID)
	if err != nil {
		return product.Product{}, err
	}
	return result, nil
}