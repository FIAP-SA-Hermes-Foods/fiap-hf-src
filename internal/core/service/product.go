package service

import "fiap-hf-src/internal/core/domain/entity"

type ProductService interface {
	SaveProduct(order entity.Product) (*entity.Product, error)
}

type productService struct {
	Product *entity.Product
}

func NewProductService(product *entity.Product) ProductService {
	if product == nil {
		return productService{Product: new(entity.Product)}
	}
	return productService{Product: product}
}

func (p productService) SaveProduct(product entity.Product) (*entity.Product, error) {

	if err := product.Category.Validate(); err != nil {
		return nil, err
	}

	return &product, nil
}
