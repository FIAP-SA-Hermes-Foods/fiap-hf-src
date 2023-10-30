package application

import (
	"context"
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	httpHF "fiap-hf-src/internal/core/domain/http"
	"fiap-hf-src/internal/core/domain/repository"
	"fiap-hf-src/internal/core/domain/valueObject"
	"fiap-hf-src/internal/core/service"
	"fmt"
	"log"
)

type HermesFoodsApp interface {
	// Client Methods

	SaveClient(client entity.Client) (*entity.OutputClient, error)
	GetClientByCPF(cpf string) (*entity.OutputClient, error)
	GetClientByID(id int64) (*entity.OutputClient, error)

	// Order Methods
	SaveOrder(order entity.Order) (*entity.OutputOrder, error)
	GetOrderByID(id int64) (*entity.OutputOrder, error)
	GetOrders() ([]entity.OutputOrder, error)
}

type hermesFoodsApp struct {
	Ctx           context.Context
	paymentAPI    httpHF.PaymentAPI
	clientRepo    repository.ClientRepository
	clientService service.ClientService
	orderRepo     repository.OrderRepository
	orderService  service.OrderService
}

func NewHermesFoodsApp(ctx context.Context, paymentAPI httpHF.PaymentAPI, clientRepo repository.ClientRepository, orderRepo repository.OrderRepository, clientService service.ClientService, orderService service.OrderService) HermesFoodsApp {
	return hermesFoodsApp{
		Ctx:           ctx,
		paymentAPI:    paymentAPI,
		clientRepo:    clientRepo,
		clientService: clientService,
		orderRepo:     orderRepo,
		orderService:  orderService,
	}
}

func (h hermesFoodsApp) GetClientByID(id int64) (*entity.OutputClient, error) {
	if err := h.GetClientByIDService(id); err != nil {
		return nil, err
	}

	c, err := h.GetClientByIDRepository(id)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &entity.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	return out, err
}

func (h hermesFoodsApp) GetClientByCPF(cpf string) (*entity.OutputClient, error) {
	if err := h.GetClientByCPFService(cpf); err != nil {
		return nil, err
	}

	c, err := h.GetClientByCPFRepository(cpf)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &entity.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	return out, err
}

func (h hermesFoodsApp) SaveClient(client entity.Client) (*entity.OutputClient, error) {
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

	cRepo, err := h.SaveClientRepository(*c)

	if err != nil {
		return nil, err
	}

	out := &entity.OutputClient{
		ID:        cRepo.ID,
		Name:      cRepo.Name,
		CPF:       cRepo.CPF.Value,
		Email:     cRepo.Email,
		CreatedAt: cRepo.CreatedAt.Format(),
	}

	return out, nil
}

func (h hermesFoodsApp) GetOrders() ([]entity.OutputOrder, error) {

	orderList := make([]entity.OutputOrder, 0)

	orders, err := h.GetOrdersRepository()

	if err != nil {
		return nil, err
	}

	for i := range orders {

		client, err := h.GetClientByID(orders[i].ID)

		if err != nil {
			return nil, err
		}

		order := entity.OutputOrder{
			ID: orders[i].ID,
			Client: entity.OutputClient{
				ID:        client.ID,
				Name:      client.Name,
				CPF:       client.CPF,
				Email:     client.Email,
				CreatedAt: client.CreatedAt,
			},
			VoucherID:        orders[i].VoucherID,
			Status:           orders[i].Status.Value,
			VerificationCode: orders[i].VerificationCode.Value,
			CreatedAt:        orders[i].CreatedAt.Format(),
		}

		orderList = append(orderList, order)
	}

	return orderList, nil
}

func (h hermesFoodsApp) GetOrderByID(id int64) (*entity.OutputOrder, error) {
	if err := h.orderService.GetOrderByID(id); err != nil {
		return nil, err
	}

	o, err := h.GetOrderByIDRepository(id)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outClient, err := h.GetClientByID(o.ClientID)

	if err != nil {
		return nil, err
	}

	if outClient == nil {
		return nil, errors.New("client is null")
	}

	out := &entity.OutputOrder{
		ID:               o.ID,
		Client:           *outClient,
		VoucherID:        o.VoucherID,
		Status:           o.Status.Value,
		VerificationCode: o.VerificationCode.Value,
		CreatedAt:        o.CreatedAt.Format(),
	}

	return out, nil
}

func (h hermesFoodsApp) SaveOrder(order entity.Order) (*entity.OutputOrder, error) {
	if err := h.GetClientByIDService(order.ClientID); err != nil {
		return nil, err
	}

	c, err := h.GetClientByIDRepository(order.ClientID)

	if err != nil {
		return nil, err
	}

	inputDoPaymentAPI := entity.InputPaymentAPI{
		Price: 0.0,
		Client: entity.Client{
			ID:   c.ID,
			Name: c.Name,
			CPF: valueObject.Cpf{
				Value: c.CPF.Value,
			},
			Email: c.Email,
			CreatedAt: valueObject.CreatedAt{
				Value: c.CreatedAt.Value,
			},
		},
	}

	out, err := h.DoPaymentAPI(h.Ctx, inputDoPaymentAPI)

	if err != nil {
		return nil, err
	}

	if out.Error != nil {
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	log.Printf("output mercado pago api: %s\n", out.MarshalString())

	order.Status.Value = out.PaymentStatus

	o, err := h.SaveOrderService(order)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, errors.New("is not possible to save order because it's null")
	}

	oRepo, err := h.SaveOrderRepository(order)

	if err != nil {
		return nil, err
	}

	outClient := entity.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	outOrder := &entity.OutputOrder{
		ID:               oRepo.ID,
		Client:           outClient,
		VoucherID:        oRepo.VoucherID,
		Status:           oRepo.Status.Value,
		VerificationCode: oRepo.VerificationCode.Value,
		CreatedAt:        oRepo.CreatedAt.Format(),
	}

	return outOrder, nil
}

func (h hermesFoodsApp) DoPaymentAPI(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error) {
	return h.paymentAPI.DoPayment(ctx, input)
}

// Client implementation Call

func (h hermesFoodsApp) GetClientByCPFService(cpf string) error {
	return h.clientService.GetClientByCPF(cpf)
}

func (h hermesFoodsApp) GetClientByCPFRepository(cpf string) (*entity.Client, error) {
	return h.clientRepo.GetClientByCPF(cpf)
}

func (h hermesFoodsApp) GetClientByIDService(id int64) error {
	return h.clientService.GetClientByID(id)
}

func (h hermesFoodsApp) GetClientByIDRepository(id int64) (*entity.Client, error) {
	return h.clientRepo.GetClientByID(id)
}

func (h hermesFoodsApp) SaveClientService(client entity.Client) (*entity.Client, error) {
	return h.clientService.SaveClient(client)
}

func (h hermesFoodsApp) SaveClientRepository(client entity.Client) (*entity.Client, error) {
	return h.clientRepo.SaveClient(client)
}

// Order implementation Call

func (h hermesFoodsApp) GetOrdersRepository() ([]entity.Order, error) {
	return h.orderRepo.GetOrders()
}

func (h hermesFoodsApp) GetOrderByIDRepository(id int64) (*entity.Order, error) {
	return h.orderRepo.GetOrderByID(id)
}

func (h hermesFoodsApp) GetOrderByIDService(id int64) error {
	return h.orderService.GetOrderByID(id)
}

func (h hermesFoodsApp) SaveOrderRepository(order entity.Order) (*entity.Order, error) {
	return h.orderRepo.SaveOrder(order)
}

func (h hermesFoodsApp) SaveOrderService(order entity.Order) (*entity.Order, error) {
	return h.orderService.SaveOrder(order)
}
