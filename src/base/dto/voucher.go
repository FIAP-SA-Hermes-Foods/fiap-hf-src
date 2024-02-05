package dto

import (
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/operation/presenter/common"
	"time"
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
	expirationTime, _ := time.Parse("02-01-2006 15:04:05", r.ExpiresAt)
	createdAt, _ := time.Parse("02-01-2006 15:04:05", r.CreatedAt)

	return entity.Voucher{
		Code:       r.Code,
		Porcentage: r.Porcentage,
		ExpiresAt: common.ExpiresAt{
			Value: &expirationTime,
		},
		CreatedAt: common.CreatedAt{
			Value: createdAt,
		},
	}
}
