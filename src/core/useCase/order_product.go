package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.OrderProductUseCase = (*orderProductUseCase)(nil)

type orderProductUseCase struct {
	gateway interfaces.OrderProductGateway
}

func NewOrderProductUseCase(gateway interfaces.OrderProductGateway) *orderProductUseCase {
	return &orderProductUseCase{gateway: gateway}
}

func (o *orderProductUseCase) GetAllOrderProductByOrderID(id int64) ([]dto.OutputOrderProduct, error) {
	if id < 1 {
		return nil, errors.New("the order id is not valid for consult")
	}

	return o.gateway.GetAllOrderProductByOrderID(id)
}

func (o *orderProductUseCase) SaveOrderProduct(order dto.RequestOrderProduct) (*dto.OutputOrderProduct, error) {
	if order.OrderID < 1 {
		return nil, errors.New("the order id is not valid for consult")
	}

	return o.gateway.SaveOrderProduct(order)
}

func (o *orderProductUseCase) GetAllOrderProduct() ([]dto.OutputOrderProduct, error) {
	return o.gateway.GetAllOrderProduct()
}
