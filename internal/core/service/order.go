package service

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
)

type OrderService interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
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

func (c orderService) SaveOrder(order entity.Order) (*entity.Order, error) {

	if err := order.Status.Validate(); err != nil {
		return nil, err
	}

	order.VerificationCode.Generate()

	if err := order.VerificationCode.Validate(); err != nil {
		order.VerificationCode.Generate()
	}

	return &order, nil
}

func (c orderService) GetClientByCPF(cpf string) error {
	if len(cpf) == 0 {
		return errors.New("the cpf is not valid for consult")
	}
	return nil
}
