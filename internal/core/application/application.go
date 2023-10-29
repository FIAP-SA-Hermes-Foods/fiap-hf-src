package application

import (
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/repository"
	"fiap-hf-src/internal/core/service"
)

type HermesFoodsApp interface {
	// Client Methods

	SaveClient(client entity.Client) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)

	// Order Methods
	SaveOrder(order entity.Order) (*entity.Order, error)
}

type hermesFoodsApp struct {
	clientRepo    repository.ClientRepository
	clientService service.ClientService
	orderRepo     repository.OrderRepository
	orderService  service.OrderService
}

func NewHermesFoodsApp(clientRepo repository.ClientRepository, orderRepo repository.OrderRepository, clientService service.ClientService, orderService service.OrderService) HermesFoodsApp {
	return hermesFoodsApp{
		clientRepo:    clientRepo,
		clientService: clientService,
		orderRepo:     orderRepo,
		orderService:  orderService,
	}
}

func (h hermesFoodsApp) GetClientByCPF(cpf string) (*entity.Client, error) {
	if err := h.GetClientByCPFService(cpf); err != nil {
		return nil, err
	}

	client, err := h.GetClientByCPFRepository(cpf)

	if err != nil {
		return nil, err
	}

	return client, err
}

func (h hermesFoodsApp) SaveClient(client entity.Client) (*entity.Client, error) {
	clientWithCpf, err := h.GetClientByCPF(client.CPF.Value)

	if err != nil {
		return nil, err
	}

	if clientWithCpf != nil {
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

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

func (h hermesFoodsApp) SaveOrder(order entity.Order) (*entity.Order, error) {
	o, err := h.SaveOrderService(order)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, errors.New("is not possible to save order because it's null")
	}

	returnORepo, err := h.SaveOrderRepository(order)
	if err != nil {
		return nil, err
	}

	return returnORepo, nil
}

func (h hermesFoodsApp) GetClientByCPFService(cpf string) error {
	return h.clientService.GetClientByCPF(cpf)
}

func (h hermesFoodsApp) GetClientByCPFRepository(cpf string) (*entity.Client, error) {
	return h.clientRepo.GetClientByCPF(cpf)
}

func (h hermesFoodsApp) SaveClientService(client entity.Client) (*entity.Client, error) {
	return h.clientService.SaveClient(client)
}

func (h hermesFoodsApp) SaveClientRepository(client entity.Client) (*entity.Client, error) {
	return h.clientRepo.SaveClient(client)
}

func (h hermesFoodsApp) SaveOrderRepository(order entity.Order) (*entity.Order, error) {
	return h.orderRepo.SaveOrder(order)
}

func (h hermesFoodsApp) SaveOrderService(order entity.Order) (*entity.Order, error) {
	return h.orderService.SaveOrder(order)
}
