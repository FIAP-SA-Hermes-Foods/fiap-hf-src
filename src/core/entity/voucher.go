package entity

import (
	com "fiap-hf-src/src/operation/presenter/common"
)

type (
	Voucher struct {
		ID         int64         `json:"id,omitempty"`
		Code       string        `json:"code,omitempty"`
		Porcentage int64         `json:"porcentage,omitempty"`
		CreatedAt  com.CreatedAt `json:"createdAt,omitempty"`
		ExpiresAt  com.ExpiresAt `json:"expiresAt,omitempty"`
	}
)
