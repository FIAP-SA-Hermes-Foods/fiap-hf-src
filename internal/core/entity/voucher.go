package entity

import (
	"encoding/json"
	com "fiap-hf-src/internal/core/entity/common"
	"log"
)

type (
	Voucher struct {
		ID         int64         `json:"id,omitempty"`
		Code       string        `json:"code,omitempty"`
		Porcentage int64         `json:"porcentage,omitempty"`
		CreatedAt  com.CreatedAt `json:"createdAt,omitempty"`
		ExpiresAt  com.ExpiresAt `json:"expiresAt,omitempty"`
	}

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

func (c Voucher) MarshalString() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (r RequestVoucher) MarshalString() string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (o OutputVoucher) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
