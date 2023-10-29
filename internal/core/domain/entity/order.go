package entity

import (
	"encoding/json"
	"fiap-hf-src/internal/core/domain/valueObject"
	"log"
)

type (
	Order struct {
		ID               int64                        `json:"id,omitempty"`
		ClientID         int64                        `json:"clientId,omitempty"`
		VoucherID        int64                        `json:"voucherId,omitempty"`
		Status           valueObject.Status           `json:"status,omitempty"`
		VerificationCode valueObject.VerificationCode `json:"verificationCode,omitempty"`
		CreatedAt        valueObject.CreatedAt        ` json:"createdAt,omitempty"`
	}

	RequestOrder struct {
		ID               int64  `json:"id,omitempty"`
		ClientID         int64  `json:"clientId,omitempty"`
		VoucherID        int64  `json:"voucherId,omitempty"`
		Status           string `json:"status,omitempty"`
		VerificationCode string `json:"verificationCode,omitempty"`
		CreatedAt        string ` json:"createdAt,omitempty"`
	}
)

func (o Order) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (o RequestOrder) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
