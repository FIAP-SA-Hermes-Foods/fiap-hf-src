package entity

import (
	com "fiap-hf-src/src/core/entity/common"
)

type (
	Order struct {
		ID               int64                `json:"id,omitempty"`
		ClientID         int64                `json:"clientId,omitempty"`
		VoucherID        *int64               `json:"voucherId,omitempty"`
		Items            []OrderItems         `json:"items,omitempty"`
		Status           com.Status           `json:"status,omitempty"`
		VerificationCode com.VerificationCode `json:"verificationCode,omitempty"`
		CreatedAt        com.CreatedAt        `json:"createdAt,omitempty"`
	}

	OrderItems struct {
		ProductID int64 `json:"productID,omitempty"`
		Quantity  int64 `json:"quantity,omitempty"`
	}
)
