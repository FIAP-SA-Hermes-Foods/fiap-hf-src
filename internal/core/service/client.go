package service

import (
	"fiap-hf-src/internal/core/domain/entity"
)

type ClientService interface {
	SaveClient(client entity.Client) (*entity.Client, error)
}

type clientService struct {
}

func NewClientService() ClientService {
	return clientService{}
}

func (c clientService) SaveClient(client entity.Client) (*entity.Client, error) {

	if err := client.CPF.Validate(); err != nil {
		return nil, err
	}

	return &client, nil
}
