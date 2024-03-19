package interfaces

import (
	"context"
	"fiap-hf-src/src/base/dto"
)

type PaymentAPI interface {
	DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error)
}

type UserAuth interface {
	Auth(in dto.UserInput) (*dto.UserOutput, error)
}
