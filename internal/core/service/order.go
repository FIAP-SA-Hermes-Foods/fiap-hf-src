package service

import (
	"errors"
	"fiap-hf-src/internal/core/entity"
)

type OrderService interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	GetOrderByID(id int64) error
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
}

type orderService struct {
	Order *entity.Order
}

func NewOrderService(order *entity.Order) OrderService {
	if order == nil {
		return orderService{Order: new(entity.Order)}
	}
	return orderService{Order: order}
}

func (o orderService) SaveOrder(order entity.Order) (*entity.Order, error) {

	if err := order.Status.Validate(); err != nil {
		return nil, err
	}

	if len(order.VerificationCode.Value) == 0 {
		order.VerificationCode.Generate()
	}

	if err := order.VerificationCode.Validate(); err != nil {
		order.VerificationCode.Generate()
	}

	return &order, nil
}

func (o orderService) UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	if err := order.Status.Validate(); err != nil {
		return nil, err
	}

	if err := order.VerificationCode.Validate(); err != nil {
		order.VerificationCode.Generate()
	}

	return &order, nil
}

func (c orderService) GetOrderByID(id int64) error {
	if id < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}
