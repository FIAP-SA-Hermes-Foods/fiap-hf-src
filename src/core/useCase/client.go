package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.ClientUseCase = (*clientUseCase)(nil)

type clientUseCase struct {
	gateway interfaces.ClientGateway
}

func NewClientUseCase(gateway interfaces.ClientGateway) clientUseCase {
	return clientUseCase{gateway: gateway}
}

func (c clientUseCase) SaveClient(reqClient dto.RequestClient) (*dto.OutputClient, error) {
	client := reqClient.Client()

	if err := client.CPF.Validate(); err != nil {
		return nil, err
	}

	out, err := c.gateway.SaveClient(reqClient)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c clientUseCase) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	if len(cpf) == 0 {
		return nil, errors.New("the cpf is not valid for consult")
	}

	out, err := c.gateway.GetClientByCPF(cpf)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c clientUseCase) GetClientByID(id int64) (*dto.OutputClient, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	out, err := c.gateway.GetClientByID(id)

	if err != nil {
		return nil, err
	}

	return out, err
}
