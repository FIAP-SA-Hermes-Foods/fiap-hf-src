package dto

import (
	"fiap-hf-src/src/core/entity"
)

type (
	InputPaymentAPI struct {
		Price  float64       `json:"price,omitempty"`
		Client entity.Client `json:"client,omitempty"`
	}
)

type (
	OutputPaymentAPI struct {
		PaymentStatus string                 `json:"paymentStatus,omitempty"`
		HTTPStatus    int                    `json:"httpStatus,omitempty"`
		Error         *OutputPaymentAPIError `json:"error,omitempty"`
	}

	OutputPaymentAPIError struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
)
