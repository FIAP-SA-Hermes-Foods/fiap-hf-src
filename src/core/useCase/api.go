package useCase

import (
	"context"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.PaymentUseCase = (*paymentUseCase)(nil)

type paymentUseCase struct {
	gateway interfaces.PaymentGateway
}

func NewPaymentUseCase(gateway interfaces.PaymentGateway) *paymentUseCase {
	return &paymentUseCase{gateway: gateway}
}

func (p *paymentUseCase) DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error) {
	return p.gateway.DoPayment(ctx, input)
}
