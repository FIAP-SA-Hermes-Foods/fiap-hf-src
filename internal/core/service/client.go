package service

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
)

type ClientService interface {
	SaveClient(client entity.Client) (*entity.Client, error)
	GetClientByCPF(cpf string) error
}

type clientService struct {
	Client *entity.Client
}

func NewClientService(client *entity.Client) ClientService {
	if client == nil {
		return clientService{Client: new(entity.Client)}
	}
	return clientService{Client: client}
}

func (c clientService) SaveClient(client entity.Client) (*entity.Client, error) {
	if err := client.CPF.Validate(); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c clientService) GetClientByCPF(cpf string) error {
	if len(cpf) == 0 {
		return errors.New("the cpf is not valid for consult")
	}
	return nil
}
