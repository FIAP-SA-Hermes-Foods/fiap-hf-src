package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.OrderUseCase = (*orderUseCase)(nil)

type orderUseCase struct {
	gateway interfaces.OrderGateway
}

func NewOrderUseCase(gateway interfaces.OrderGateway) *orderUseCase {
	return &orderUseCase{gateway: gateway}
}

func (o *orderUseCase) SaveOrder(reqOrder dto.RequestOrder) (*dto.OutputOrder, error) {

	order := reqOrder.Order()

	if err := order.Status.Validate(); err != nil {
		return nil, err
	}

	if len(order.VerificationCode.Value) == 0 {
		order.VerificationCode.Generate()
	}

	if err := order.VerificationCode.Validate(); err != nil {
		order.VerificationCode.Generate()
	}

	reqOrder.VerificationCode = order.VerificationCode.Value

	out, err := o.gateway.SaveOrder(reqOrder)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *orderUseCase) UpdateOrderByID(id int64, reqOrder dto.RequestOrder) (*dto.OutputOrder, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	order := reqOrder.Order()

	if err := order.Status.Validate(); err != nil {
		return nil, err
	}

	if err := order.VerificationCode.Validate(); err != nil {
		order.VerificationCode.Generate()
	}

	reqOrder.VerificationCode = order.VerificationCode.Value

	out, err := o.gateway.UpdateOrderByID(id, reqOrder)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *orderUseCase) GetOrders() ([]dto.OutputOrder, error) {
	return o.gateway.GetOrders()
}

func (o *orderUseCase) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	out, err := o.gateway.GetOrderByID(id)

	if err != nil {
		return nil, err
	}

	return out, nil
}
