package interfaces

import "fiap-hf-src/src/core/entity"

type OrderRepository interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
	GetOrders() ([]entity.Order, error)
	GetOrderByID(id int64) (*entity.Order, error)
}

type OrderService interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	GetOrderByID(id int64) error
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
}
