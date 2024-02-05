package db

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.ClientGateway = (*clientGateway)(nil)

type clientGateway struct {
	db interfaces.ClientDB
}

func NewClientGateway(db interfaces.ClientDB) *clientGateway {
	return &clientGateway{db: db}
}

func (g *clientGateway) GetClientByID(id int64) (*dto.OutputClient, error) {

	c, err := g.db.GetClientByID(id)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}
	return out, nil
}

func (g *clientGateway) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	c, err := g.db.GetClientByCPF(cpf)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	return out, nil
}

func (g *clientGateway) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {

	c, err := g.db.SaveClient(client.Client())

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	return out, nil
}
