package entity

import (
	"encoding/json"
	com "fiap-hf-src/internal/core/entity/common"
	"log"
)

type (
	Client struct {
		ID        int64         `json:"id,omitempty"`
		Name      string        `json:"name,omitempty"`
		CPF       com.Cpf       `json:"cpf,omitempty"`
		Email     string        `json:"email,omitempty"`
		CreatedAt com.CreatedAt ` json:"createdAt,omitempty"`
	}

	RequestClient struct {
		ID        int64  `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		CreatedAt string `json:"createdAt,omitempty"`
	}

	OutputClient struct {
		ID        int64  `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		CreatedAt string `json:"createdAt,omitempty"`
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

func (r RequestClient) MarshalString() string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}

func (o OutputClient) MarshalString() string {
	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("error in MarshalString client %v", err)
		return ""
	}

	return string(b)
}
