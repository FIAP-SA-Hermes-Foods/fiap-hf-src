package dto

import (
	"fiap-hf-src/src/core/entity"
	com "fiap-hf-src/src/operation/presenter/common"
)

type (
	RequestProduct struct {
		ID            int64   `json:"id,omitempty"`
		Name          string  `json:"name,omitempty"`
		Category      string  `json:"category,omitempty"`
		Image         string  `json:"image,omitempty"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price,omitempty"`
		CreatedAt     string  `json:"createdAt,omitempty"`
		DeactivatedAt string  `json:"deactivatedAt,omitempty"`
	}

	OutputProduct struct {
		ID            int64   `json:"id,omitempty"`
		Name          string  `json:"name,omitempty"`
		Category      string  `json:"category,omitempty"`
		Image         string  `json:"image,omitempty"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price,omitempty"`
		CreatedAt     string  `json:"createdAt,omitempty"`
		DeactivatedAt string  `json:"deactivatedAt,omitempty"`
	}
)

func (r RequestProduct) Product() entity.Product {
	product := entity.Product{
		Name: r.Name,
		Category: com.Category{
			Value: r.Category,
		},
		Image:       r.Image,
		Description: r.Description,
		Price:       r.Price,
	}

	return product
}
