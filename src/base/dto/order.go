package dto

import (
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/core/entity/common"
)

type (
	RequestOrder struct {
		ID               int64               `json:"id,omitempty"`
		ClientID         int64               `json:"clientId,omitempty"`
		VoucherID        *int64              `json:"voucherId,omitempty"`
		Items            []entity.OrderItems `json:"items,omitempty"`
		Status           string              `json:"status,omitempty"`
		VerificationCode string              `json:"verificationCode,omitempty"`
		CreatedAt        string              `json:"createdAt,omitempty"`
	}

	OutputOrder struct {
		ID               int64                `json:"id,omitempty"`
		Client           OutputClient         `json:"client,omitempty"`
		Products         []entity.ProductItem `json:"products,omitempty"`
		VoucherID        *int64               `json:"voucherId,omitempty"`
		Status           string               `json:"status,omitempty"`
		VerificationCode string               `json:"verificationCode,omitempty"`
		CreatedAt        string               `json:"createdAt,omitempty"`
		TotalPrice       float64              `json:"totalPrice"`
	}
)

func (r RequestOrder) Order() entity.Order {
	return entity.Order{
		ClientID:  r.ClientID,
		VoucherID: r.VoucherID,
		Items:     r.Items,
		Status: common.Status{
			Value: r.Status,
		},
	}
}
