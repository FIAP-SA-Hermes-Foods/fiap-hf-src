package http

import (
	"context"
	"hermes-foods/internal/core/domain/entity"
)

type PaymentAPI interface {
	DoPayment(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error)
}
