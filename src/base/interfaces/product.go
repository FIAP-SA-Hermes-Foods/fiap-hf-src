package interfaces

import "fiap-hf-src/src/core/entity"

type ProductRepository interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByCategory(category string) ([]entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	DeleteProductByID(id int64) error
}

type ProductService interface {
	SaveProduct(order entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByID(id int64) error
	GetProductByCategory(category string) error
	DeleteProductByID(id int64) error
}
