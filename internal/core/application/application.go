package application

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/repository"
	"fiap-hf-src/internal/core/service"
)

type HermesFoodsApp interface {
	SaveClient(client entity.Client) (*entity.Client, error)
}

type hermesFoodsApp struct {
	clientRepo    repository.ClientRepository
	clientService service.ClientService
}

func NewHermesFoodsApp(clientRepo repository.ClientRepository, clientService service.ClientService) HermesFoodsApp {
	return hermesFoodsApp{
		clientRepo:    clientRepo,
		clientService: clientService,
	}
}

func (h hermesFoodsApp) SaveClient(client entity.Client) (*entity.Client, error) {
	c, err := h.SaveClientService(client)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New("is not possible to save client because it's null")
	}

	returnCRepo, err := h.SaveClientRepository(*c)

	if err != nil {
		return nil, err
	}

	return returnCRepo, nil
}

func (h hermesFoodsApp) SaveClientService(client entity.Client) (*entity.Client, error) {
	return h.clientService.SaveClient(client)
}

func (h hermesFoodsApp) SaveClientRepository(client entity.Client) (*entity.Client, error) {
	return h.clientRepo.SaveClient(client)
}
