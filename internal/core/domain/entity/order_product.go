package entity

import (
	"encoding/json"
	"fiap-hf-src/internal/core/domain/valueObject"
	"log"
)

type (
	OrderProduct struct {
		ID         int64                 `json:"id,omitempty"`
		Quantity   int64                 `json:"quantity,omitempty"`
		TotalPrice float64               `json:"totalPrice,omitempty"`
		Discount   float64               `json:"discount,omitempty"`
		OrderID    int64                 `json:"orderId,omitempty"`
		ProductID  *int64                `json:"productId,omitempty"`
		CreatedAt  valueObject.CreatedAt ` json:"createdAt,omitempty"`
	}

	RequestOrderProduct struct {
		ID         int64   `json:"id,omitempty"`
		Quantity   int64   `json:"quantity,omitempty"`
		TotalPrice float64 `json:"totalPrice,omitempty"`
		Discount   float64 `json:"discount,omitempty"`
		OrderID    int64   `json:"orderId,omitempty"`
		ProductID  *int64  `json:"productId,omitempty"`
		CreatedAt  string  `json:"createdAt,omitempty"`
	}

	OutputOrderProduct struct {
		ID         int64   `json:"id,omitempty"`
		Quantity   int64   `json:"quantity,omitempty"`
		TotalPrice float64 `json:"totalPrice,omitempty"`
		Discount   float64 `json:"discount,omitempty"`
		OrderID    int64   `json:"orderId,omitempty"`
		ProductID  *int64  `json:"productId,omitempty"`
		CreatedAt  string  `json:"createdAt,omitempty"`
	}
)

func (o OrderProduct) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}
	return string(b)
}

func (o RequestOrderProduct) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}
	return string(b)
}

func (o OutputOrderProduct) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}
	return string(b)
}
