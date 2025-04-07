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

func (s *service) UpdateProductByID(productid uint, newData product.Product) (product.Product, error) {
	existingProduct, err := s.model.GetProductByID(productid)
	if err != nil {
		return product.Product{}, err
	}
	if newData.Nama != "" {
		existingProduct.Nama = newData.Nama
	}
	if newData.Brand != "" {
		existingProduct.Brand = newData.Brand
	}
	if newData.Category != "" {
		existingProduct.Category = newData.Category
	}
	if newData.Price != 0 {
		existingProduct.Price = newData.Price
	}
	if newData.Amount != 0 {
		existingProduct.Amount = newData.Amount
	}
	if newData.Description != "" {
		existingProduct.Description = newData.Description
	}
	if newData.Image != "" {
		existingProduct.Image = newData.Image
	}
	result, err := s.model.UpdateProductByID(productid, existingProduct)
	if err != nil {
		return product.Product{}, err
	}
	return result, nil
}