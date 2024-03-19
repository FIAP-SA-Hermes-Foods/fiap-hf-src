package http

import (
	"context"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.PaymentGateway = (*mercadoPagoGateway)(nil)

type mercadoPagoGateway struct {
	api  interfaces.PaymentAPI
	auth interfaces.UserAuth
}

func NewMercadoPagoAPI(api interfaces.PaymentAPI) *mercadoPagoGateway {
	return &mercadoPagoGateway{api: api}
}

func (m *mercadoPagoGateway) DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error) {
	return m.api.DoPayment(ctx, input)
}
