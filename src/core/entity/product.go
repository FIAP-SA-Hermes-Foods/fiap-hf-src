package entity

import (
	com "fiap-hf-src/src/operation/presenter/common"
)

type (
	Product struct {
		ID            int64             `json:"id,omitempty"`
		Name          string            `json:"name,omitempty"`
		Category      com.Category      `json:"category,omitempty"`
		Image         string            `json:"image,omitempty"`
		Description   string            `json:"description,omitempty"`
		Price         float64           `json:"price,omitempty"`
		CreatedAt     com.CreatedAt     `json:"createdAt,omitempty"`
		DeactivatedAt com.DeactivatedAt `json:"deactivatedAt,omitempty"`
	}

	ProductItem struct {
		ID            int64   `json:"id,omitempty"`
		Name          string  `json:"name,omitempty"`
		Category      string  `json:"category,omitempty"`
		Image         string  `json:"image,omitempty"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price,omitempty"`
		Quantity      int64   `json:"quantity,omitempty"`
		CreatedAt     string  `json:"createdAt,omitempty"`
		DeactivatedAt string  `json:"deactivatedAt,omitempty"`
	}
)
