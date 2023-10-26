package entity

import (
	"encoding/json"
	"log"
)

type InputPaymentAPI struct {
	Price  float64 `json:"price,omitempty"`
	Client Client  `json:"client,omitempty"`
}

func (i InputPaymentAPI) MarshalString() string {
	b, err := json.Marshal(i)

	if err != nil {
		log.Printf("error in marshal InputPaymentAPI: %v", err)
		return ""
	}
	return string(b)
}

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

func (o OutputPaymentAPI) MarshalString() string {
	b, err := json.Marshal(o)

	if err != nil {
		log.Printf("error in marshal OutputPaymentAPI: %v", err)
		return ""
	}
	return string(b)
}
