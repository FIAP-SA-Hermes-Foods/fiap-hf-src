package repository

import "fiap-hf-src/internal/core/domain/entity"

type ClientRepository interface {
	GetClientByID(id int64) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type OrderRepository interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
	GetOrders() ([]entity.Order, error)
	GetOrderByID(id int64) (*entity.Order, error)
}

type ProductRepository interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByCategory(category string) ([]entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	DeleteProductByID(id int64) error
}

type OrderProductRepository interface {
	GetAllOrderProduct() ([]entity.OrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error)
}
