package dto

import (
	"fiap-hf-src/src/core/entity"
)

type (
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

func (r RequestOrderProduct) OrderProduct() entity.OrderProduct {
	return entity.OrderProduct{
		Quantity:   r.Quantity,
		TotalPrice: r.TotalPrice,
		Discount:   r.Discount,
		OrderID:    r.OrderID,
		ProductID:  r.ProductID,
	}
}
