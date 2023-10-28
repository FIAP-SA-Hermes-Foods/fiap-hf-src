package entity

import (
	"encoding/json"
	"fiap-hf-src/internal/core/domain/valueObject"
	"log"
)

type (
	Client struct {
		ID    int64           `json:"id,omitempty"`
		Name  string          `json:"name,omitempty"`
		CPF   valueObject.Cpf `json:"cpf,omitempty"`
		Email string          `json:"email,omitempty"`
	}

	RequestClient struct {
		ID    int64  `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		CPF   string `json:"cpf,omitempty"`
		Email string `json:"email,omitempty"`
	}
)

func (c Client) MarshalString() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
