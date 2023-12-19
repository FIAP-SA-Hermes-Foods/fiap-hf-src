package entity

import (
	"encoding/json"
	com "fiap-hf-src/internal/core/entity/common"
	"log"
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

	RequestOrder struct {
		ID               int64        `json:"id,omitempty"`
		ClientID         int64        `json:"clientId,omitempty"`
		VoucherID        *int64       `json:"voucherId,omitempty"`
		Items            []OrderItems `json:"items,omitempty"`
		Status           string       `json:"status,omitempty"`
		VerificationCode string       `json:"verificationCode,omitempty"`
		CreatedAt        string       `json:"createdAt,omitempty"`
	}

	OutputOrder struct {
		ID               int64         `json:"id,omitempty"`
		Client           OutputClient  `json:"client,omitempty"`
		Products         []ProductItem `json:"products,omitempty"`
		VoucherID        *int64        `json:"voucherId,omitempty"`
		Status           string        `json:"status,omitempty"`
		VerificationCode string        `json:"verificationCode,omitempty"`
		CreatedAt        string        `json:"createdAt,omitempty"`
		TotalPrice       float64       `json:"totalPrice"`
	}
)

type OrderItems struct {
	ProductID int64 `json:"productID,omitempty"`
	Quantity  int64 `json:"quantity,omitempty"`
}

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

func (o OutputOrder) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
