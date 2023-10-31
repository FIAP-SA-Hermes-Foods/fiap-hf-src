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
	UpdateOrderByID(id int64, order entity.Order) (*entity.OutputOrder, error)
	GetOrders() ([]entity.OutputOrder, error)

	// Product Methods

	SaveProduct(product entity.Product) (*entity.OutputProduct, error)
	GetProductByCategory(category string) ([]entity.OutputProduct, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.OutputProduct, error)
	DeleteProductByID(id int64) error
}

type hermesFoodsApp struct {
	Ctx            context.Context
	paymentAPI     httpHF.PaymentAPI
	clientRepo     repository.ClientRepository
	clientService  service.ClientService
	orderRepo      repository.OrderRepository
	orderService   service.OrderService
	productRepo    repository.ProductRepository
	productService service.ProductService
}

func NewHermesFoodsApp(ctx context.Context, paymentAPI httpHF.PaymentAPI, clientRepo repository.ClientRepository, orderRepo repository.OrderRepository, productRepo repository.ProductRepository, clientService service.ClientService, orderService service.OrderService, productService service.ProductService) HermesFoodsApp {
	return hermesFoodsApp{
		Ctx:            ctx,
		paymentAPI:     paymentAPI,
		clientRepo:     clientRepo,
		clientService:  clientService,
		orderRepo:      orderRepo,
		orderService:   orderService,
		productRepo:    productRepo,
		productService: productService,
	}
}

func (app hermesFoodsApp) GetClientByID(id int64) (*entity.OutputClient, error) {
	if err := app.GetClientByIDService(id); err != nil {
		return nil, err
	}

	c, err := app.GetClientByIDRepository(id)

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

func (app hermesFoodsApp) GetClientByCPF(cpf string) (*entity.OutputClient, error) {
	if err := app.GetClientByCPFService(cpf); err != nil {
		return nil, err
	}

	c, err := app.GetClientByCPFRepository(cpf)

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

func (app hermesFoodsApp) SaveClient(client entity.Client) (*entity.OutputClient, error) {
	clientWithCpf, err := app.GetClientByCPF(client.CPF.Value)

	if err != nil {
		return nil, err
	}

	if clientWithCpf != nil {
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

	c, err := app.SaveClientService(client)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New("is not possible to save client because it's null")
	}

	cRepo, err := app.SaveClientRepository(*c)

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

func (app hermesFoodsApp) UpdateOrderByID(id int64, order entity.Order) (*entity.OutputOrder, error) {
	oSvc, err := app.UpdateOrderByIDService(id, order)

	if err != nil {
		return nil, err
	}

	if oSvc == nil {
		return nil, errors.New("order is null, is not possible to proceed with update order")
	}

	oRepo, err := app.UpdateOrderByIDRepository(id, order)

	if err != nil {
		return nil, err
	}

	client, err := app.GetClientByID(oRepo.ClientID)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, errors.New("client is null, is not possible to proceed with update order")
	}

	out := &entity.OutputOrder{
		ID:               oRepo.ID,
		Client:           *client,
		VoucherID:        oRepo.VoucherID,
		Status:           oRepo.Status.Value,
		VerificationCode: oRepo.VerificationCode.Value,
		CreatedAt:        oRepo.CreatedAt.Format(),
	}

	return out, nil
}

func (app hermesFoodsApp) GetOrders() ([]entity.OutputOrder, error) {
	orderList := make([]entity.OutputOrder, 0)

	orders, err := app.GetOrdersRepository()

	if err != nil {
		return nil, err
	}

	for i := range orders {

		client, err := app.GetClientByID(orders[i].ID)

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

func (app hermesFoodsApp) GetOrderByID(id int64) (*entity.OutputOrder, error) {
	if err := app.orderService.GetOrderByID(id); err != nil {
		return nil, err
	}

	o, err := app.GetOrderByIDRepository(id)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outClient, err := app.GetClientByID(o.ClientID)

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

func (app hermesFoodsApp) GetProductByCategory(category string) ([]entity.OutputProduct, error) {
	productList := make([]entity.OutputProduct, 0)

	if err := app.productService.GetProductByCategory(category); err != nil {
		return nil, err
	}

	products, err := app.GetProductByCategoryRepository(category)

	if err != nil {
		return nil, err
	}

	if products == nil {
		return nil, nil
	}

	for i := range products {
		product := entity.OutputProduct{
			ID:            products[i].ID,
			Name:          products[i].Name,
			Category:      products[i].Category.Value,
			Image:         products[i].Image,
			Description:   products[i].Description,
			Price:         products[i].Price,
			CreatedAt:     products[i].CreatedAt.Format(),
			DeactivatedAt: products[i].CreatedAt.Format(),
		}
		productList = append(productList, product)
	}

	return productList, nil
}

func (app hermesFoodsApp) SaveOrder(order entity.Order) (*entity.OutputOrder, error) {
	if err := app.GetClientByIDService(order.ClientID); err != nil {
		return nil, err
	}

	c, err := app.GetClientByIDRepository(order.ClientID)

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

	out, err := app.DoPaymentAPI(app.Ctx, inputDoPaymentAPI)

	if err != nil {
		return nil, err
	}

	if out.Error != nil {
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	log.Printf("output mercado pago api: %s\n", out.MarshalString())

	order.Status.Value = out.PaymentStatus

	o, err := app.SaveOrderService(order)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, errors.New("is not possible to save order because it's null")
	}

	oRepo, err := app.SaveOrderRepository(order)

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

func (app hermesFoodsApp) SaveProduct(product entity.Product) (*entity.OutputProduct, error) {
	p, err := app.SaveProductService(product)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, errors.New("is not possible to save product because it's null")
	}

	pRepo, err := app.SaveProductRepository(*p)

	if err != nil {
		return nil, err
	}

	out := &entity.OutputProduct{
		ID:            pRepo.ID,
		Name:          pRepo.Name,
		Category:      pRepo.Category.Value,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt.Format(),
		DeactivatedAt: pRepo.DeactivatedAt.Format(),
	}

	return out, nil
}

func (app hermesFoodsApp) UpdateProductByID(id int64, product entity.Product) (*entity.OutputProduct, error) {
	if err := app.GetProductByIDService(id); err != nil {
		return nil, err
	}

	pByID, err := app.GetProductByIDRepository(id)

	if err != nil {
		return nil, err
	}

	if pByID == nil {
		return nil, errors.New("was not found any product with this id")
	}

	p, err := app.UpdateProductByIDService(id, product)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, errors.New("is not possible to save product because it's null")
	}

	pRepo, err := app.UpdateProductByIDRepository(id, *p)

	if err != nil {
		return nil, err
	}

	out := &entity.OutputProduct{
		ID:            pRepo.ID,
		Name:          pRepo.Name,
		Category:      pRepo.Category.Value,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt.Format(),
		DeactivatedAt: pRepo.DeactivatedAt.Format(),
	}
	return out, nil
}

func (app hermesFoodsApp) DeleteProductByID(id int64) error {
	if err := app.GetProductByIDService(id); err != nil {
		return err
	}

	pByID, err := app.GetProductByIDRepository(id)

	if err != nil {
		return err
	}

	if pByID == nil {
		return errors.New("was not found any product with this id")
	}

	if err := app.DeleteProductByIDService(id); err != nil {
		return err
	}

	return app.DeleteProductByIDRepository(id)
}

// ============= Calling Repositories and Services ================

func (app hermesFoodsApp) DoPaymentAPI(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error) {
	return app.paymentAPI.DoPayment(ctx, input)
}

// Client implementation Call

func (app hermesFoodsApp) GetClientByCPFService(cpf string) error {
	return app.clientService.GetClientByCPF(cpf)
}

func (app hermesFoodsApp) GetClientByCPFRepository(cpf string) (*entity.Client, error) {
	return app.clientRepo.GetClientByCPF(cpf)
}

func (app hermesFoodsApp) GetClientByIDService(id int64) error {
	return app.clientService.GetClientByID(id)
}

func (app hermesFoodsApp) GetClientByIDRepository(id int64) (*entity.Client, error) {
	return app.clientRepo.GetClientByID(id)
}

func (app hermesFoodsApp) SaveClientService(client entity.Client) (*entity.Client, error) {
	return app.clientService.SaveClient(client)
}

func (app hermesFoodsApp) SaveClientRepository(client entity.Client) (*entity.Client, error) {
	return app.clientRepo.SaveClient(client)
}

// Order implementation Call

func (app hermesFoodsApp) GetOrdersRepository() ([]entity.Order, error) {
	return app.orderRepo.GetOrders()
}

func (app hermesFoodsApp) GetOrderByIDRepository(id int64) (*entity.Order, error) {
	return app.orderRepo.GetOrderByID(id)
}

func (app hermesFoodsApp) GetOrderByIDService(id int64) error {
	return app.orderService.GetOrderByID(id)
}

func (app hermesFoodsApp) SaveOrderRepository(order entity.Order) (*entity.Order, error) {
	return app.orderRepo.SaveOrder(order)
}

func (app hermesFoodsApp) SaveOrderService(order entity.Order) (*entity.Order, error) {
	return app.orderService.SaveOrder(order)
}

func (app hermesFoodsApp) UpdateOrderByIDService(id int64, order entity.Order) (*entity.Order, error) {
	return app.orderService.UpdateOrderByID(id, order)
}

func (app hermesFoodsApp) UpdateOrderByIDRepository(id int64, order entity.Order) (*entity.Order, error) {
	return app.orderRepo.UpdateOrderByID(id, order)
}

// Product implementation Call

func (app hermesFoodsApp) GetProductByIDService(id int64) error {
	return app.productService.GetProductByID(id)
}

func (app hermesFoodsApp) GetProductByCategoryService(category string) error {
	return app.productService.GetProductByCategory(category)
}

func (app hermesFoodsApp) GetProductByIDRepository(id int64) (*entity.Product, error) {
	return app.productRepo.GetProductByID(id)
}

func (app hermesFoodsApp) GetProductByCategoryRepository(category string) ([]entity.Product, error) {
	return app.productRepo.GetProductByCategory(category)
}

func (app hermesFoodsApp) SaveProductService(product entity.Product) (*entity.Product, error) {
	return app.productService.SaveProduct(product)
}

func (app hermesFoodsApp) SaveProductRepository(product entity.Product) (*entity.Product, error) {
	return app.productRepo.SaveProduct(product)
}

func (app hermesFoodsApp) UpdateProductByIDService(id int64, product entity.Product) (*entity.Product, error) {
	return app.productService.UpdateProductByID(id, product)
}

func (app hermesFoodsApp) UpdateProductByIDRepository(id int64, product entity.Product) (*entity.Product, error) {
	return app.productRepo.UpdateProductByID(id, product)
}

func (app hermesFoodsApp) DeleteProductByIDService(id int64) error {
	return app.productService.DeleteProductByID(id)
}

func (app hermesFoodsApp) DeleteProductByIDRepository(id int64) error {
	return app.productRepo.DeleteProductByID(id)
}
