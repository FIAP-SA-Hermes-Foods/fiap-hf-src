package entity

import (
	com "fiap-hf-src/src/core/entity/common"
)

type (
	Client struct {
		ID        int64         `json:"id,omitempty"`
		Name      string        `json:"name,omitempty"`
		CPF       com.Cpf       `json:"cpf,omitempty"`
		Email     string        `json:"email,omitempty"`
		CreatedAt com.CreatedAt ` json:"createdAt,omitempty"`
	}
)
