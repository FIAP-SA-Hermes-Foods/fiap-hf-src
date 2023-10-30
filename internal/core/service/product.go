package service

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
)

type ProductService interface {
	SaveProduct(order entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByID(id int64) error
	DeleteProductByID(id int64) error
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

func (p productService) UpdateProductByID(id int64, product entity.Product) (*entity.Product, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	if err := product.Category.Validate(); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p productService) GetProductByID(id int64) error {
	if id < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}

func (p productService) DeleteProductByID(id int64) error {
	if id < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}
