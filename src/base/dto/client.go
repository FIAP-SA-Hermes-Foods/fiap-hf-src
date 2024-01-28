package dto

import (
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/operation/presenter/common"
)

type (
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

func (r RequestClient) Client() entity.Client {
	return entity.Client{
		Name: r.Name,
		CPF: common.Cpf{
			Value: r.CPF,
		},
		Email: r.Email,
	}
}
