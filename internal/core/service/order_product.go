package service

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
)

type OrderProductService interface {
	GetOrderProductByOrderID(orderID int64) error
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
