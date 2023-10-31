package entity

import (
	"encoding/json"
	"fiap-hf-src/internal/core/domain/valueObject"
	"log"
)

type (
	Product struct {
		ID            int64                     `json:"id,omitempty"`
		Name          string                    `json:"name,omitempty"`
		Category      valueObject.Category      `json:"category,omitempty"`
		Image         string                    `json:"image,omitempty"`
		Description   string                    `json:"description,omitempty"`
		Price         float64                   `json:"price,omitempty"`
		CreatedAt     valueObject.CreatedAt     `json:"createdAt,omitempty"`
		DeactivatedAt valueObject.DeactivatedAt `json:"deactivatedAt,omitempty"`
	}

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

func (p Product) MarshalString() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (p RequestProduct) MarshalString() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (p OutputProduct) MarshalString() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
