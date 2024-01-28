package dto

import (
	"fiap-hf-src/src/core/entity"
)

type (
	RequestVoucher struct {
		ID         int64  `json:"id,omitempty"`
		Code       string `json:"code,omitempty"`
		Porcentage int64  `json:"porcentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}

	OutputVoucher struct {
		ID         int64  `json:"id,omitempty"`
		Code       string `json:"code,omitempty"`
		Porcentage int64  `json:"porcentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}
)

func (r RequestVoucher) Voucher() entity.Voucher {
	return entity.Voucher{
		Code:       r.Code,
		Porcentage: r.Porcentage,
	}
}
