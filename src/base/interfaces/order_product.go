package interfaces

import "fiap-hf-src/src/core/entity"

type OrderProductRepository interface {
	GetAllOrderProduct() ([]entity.OrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error)
	SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error)
}

type OrderProductService interface {
	GetOrderProductByOrderID(orderID int64) error
	SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error)
}
