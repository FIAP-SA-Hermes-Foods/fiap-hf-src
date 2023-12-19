package http

import (
	"context"
	"fiap-hf-src/internal/core/entity"
)

type PaymentAPI interface {
	DoPayment(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error)
}
