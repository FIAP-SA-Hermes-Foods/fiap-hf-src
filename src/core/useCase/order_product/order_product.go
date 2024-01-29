package service

import (
	"errors"
	"fiap-hf-src/src/core/entity"
)

type OrderProductService interface {
	GetOrderProductByOrderID(orderID int64) error
	SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error)
}

type orderProductService struct {
	OrderProduct *entity.OrderProduct
}

func NewOrderProductService(order *entity.OrderProduct) OrderProductService {
	if order == nil {
		return orderProductService{OrderProduct: new(entity.OrderProduct)}
	}
	return orderProductService{OrderProduct: order}
}

func (o orderProductService) GetOrderProductByOrderID(orderID int64) error {
	if orderID < 1 {
		return errors.New("the order id is not valid for consult")
	}
	return nil
}

func (o orderProductService) SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error) {
	if order.OrderID < 1 {
		return nil, errors.New("the order id is not valid for consult")
	}

	return &order, nil
}
