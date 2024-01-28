package entity

import (
	com "fiap-hf-src/src/core/entity/common"
)

type (
	OrderProduct struct {
		ID         int64         `json:"id,omitempty"`
		Quantity   int64         `json:"quantity,omitempty"`
		TotalPrice float64       `json:"totalPrice,omitempty"`
		Discount   float64       `json:"discount,omitempty"`
		OrderID    int64         `json:"orderId,omitempty"`
		ProductID  *int64        `json:"productId,omitempty"`
		CreatedAt  com.CreatedAt ` json:"createdAt,omitempty"`
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
