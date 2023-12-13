package application

import (
	"context"
	"errors"
	l "fiap-hf-src/infrastructure/logger"
	"fiap-hf-src/internal/core/domain/entity"
	httpHF "fiap-hf-src/internal/core/domain/http"
	"fiap-hf-src/internal/core/domain/repository"
	"fiap-hf-src/internal/core/domain/valueObject"
	"fiap-hf-src/internal/core/service"
	"fmt"
	"log"
	"strings"
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

	// Voucher Methods

	SaveVoucher(voucher entity.Voucher) (*entity.OutputVoucher, error)
	GetVoucherByID(id int64) (*entity.OutputVoucher, error)
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.OutputVoucher, error)
}

type hermesFoodsApp struct {
	Ctx                 context.Context
	paymentAPI          httpHF.PaymentAPI
	clientRepo          repository.ClientRepository
	clientService       service.ClientService
	orderRepo           repository.OrderRepository
	orderService        service.OrderService
	orderProductRepo    repository.OrderProductRepository
	orderProductService service.OrderProductService
	productRepo         repository.ProductRepository
	productService      service.ProductService
	voucherRepo         repository.VoucherRepository
	voucherService      service.VoucherService
}

func NewHermesFoodsApp(ctx context.Context, paymentAPI httpHF.PaymentAPI, clientRepo repository.ClientRepository, orderRepo repository.OrderRepository, orderProductRepo repository.OrderProductRepository, productRepo repository.ProductRepository, voucherRepo repository.VoucherRepository, clientService service.ClientService, orderService service.OrderService, orderProductService service.OrderProductService, productService service.ProductService, voucherService service.VoucherService) HermesFoodsApp {
	return hermesFoodsApp{
		Ctx:                 ctx,
		paymentAPI:          paymentAPI,
		clientRepo:          clientRepo,
		clientService:       clientService,
		orderRepo:           orderRepo,
		orderService:        orderService,
		orderProductRepo:    orderProductRepo,
		orderProductService: orderProductService,
		productRepo:         productRepo,
		productService:      productService,
		voucherRepo:         voucherRepo,
		voucherService:      voucherService,
	}
}

// ========== Client ==========

func (app hermesFoodsApp) GetClientByID(id int64) (*entity.OutputClient, error) {
	l.Infof("GetClientByID: ", " | ", id)
	if err := app.GetClientByIDService(id); err != nil {
		l.Errorf("GetClientByIDService error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByIDRepository(id)
	l.Infof("GetClientByIDRepository output: ", " | ", c.MarshalString())

	if err != nil {
		l.Errorf("GetClientByIDRepository error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByIDRepository output: ", " | ", nil)
		return nil, nil
	}

	out := &entity.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	l.Infof("GetClientByIDRepository output: ", " | ", out.MarshalString())
	return out, err
}

func (app hermesFoodsApp) GetClientByCPF(cpf string) (*entity.OutputClient, error) {
	l.Infof("GetClientByCPF: ", " | ", cpf)
	if err := app.GetClientByCPFService(cpf); err != nil {
		l.Errorf("GetClientByCPFService error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByCPFRepository(cpf)
	l.Infof("GetClientByCPFRepository output: ", " | ", c.MarshalString())

	if err != nil {
		l.Errorf("GetClientByCPFRepository error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByCPFRepository output: ", " | ", nil)
		return nil, nil
	}

	out := &entity.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	l.Infof("GetClientByCPFRepository output: ", " | ", out.MarshalString())
	return out, err
}

func (app hermesFoodsApp) SaveClient(client entity.Client) (*entity.OutputClient, error) {
	l.Infof("SaveClient: ", " | ", client.MarshalString())
	clientWithCpf, err := app.GetClientByCPF(client.CPF.Value)

	if err != nil {
		l.Errorf("SaveClient error: ", " | ", err)
		return nil, err
	}

	if clientWithCpf != nil {
		l.Infof("SaveClient output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

	c, err := app.SaveClientService(client)
	l.Infof("SaveClientService output: ", " | ", c.MarshalString())

	if err != nil {
		l.Errorf("SaveClientService error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("SaveClientService output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because it's null")
	}

	cRepo, err := app.SaveClientRepository(*c)
	l.Infof("SaveClientRepository output: ", " | ", cRepo.MarshalString())

	if err != nil {
		l.Errorf("SaveClientRepository error: ", " | ", err)
		return nil, err
	}

	out := &entity.OutputClient{
		ID:        cRepo.ID,
		Name:      cRepo.Name,
		CPF:       cRepo.CPF.Value,
		Email:     cRepo.Email,
		CreatedAt: cRepo.CreatedAt.Format(),
	}

	l.Infof("SaveClientRepository output: ", " | ", out.MarshalString())
	return out, nil
}

// ========== Order ==========

func (app hermesFoodsApp) UpdateOrderByID(id int64, order entity.Order) (*entity.OutputOrder, error) {
	l.Infof("UpdateOrderByID: ", " | ", id, " | ", order.MarshalString())
	oSvc, err := app.UpdateOrderByIDService(id, order)

	if err != nil {
		l.Errorf("UpdateOrderByIDService error: ", " | ", err)
		return nil, err
	}

	if oSvc == nil {
		l.Infof("UpdateOrderByIDService output: ", " | ", nil)
		return nil, errors.New("order is null, is not possible to proceed with update order")
	}

	oRepo, err := app.UpdateOrderByIDRepository(id, order)
	l.Infof("UpdateOrderByIDRepository output: ", " | ", oRepo.MarshalString())

	if err != nil {
		l.Errorf("UpdateOrderByIDRepository error: ", " | ", err)
		return nil, err
	}

	client, err := app.GetClientByID(oRepo.ClientID)
	l.Infof("GetClientByID output: ", " | ", client.MarshalString())

	if err != nil {
		l.Errorf("GetClientByID error: ", " | ", err)
		return nil, err
	}

	if client == nil {
		l.Infof("GetClientByID output: ", " | ", nil)
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

	l.Infof("UpdateOrderByIDRepository output: ", " | ", out.MarshalString())
	return out, nil
}

func (app hermesFoodsApp) GetOrders() ([]entity.OutputOrder, error) {
	l.Infof("GetOrders: ", " | ")
	orderList := make([]entity.OutputOrder, 0)

	orders, err := app.GetOrdersRepository()
	//l.Infof("GetOrdersRepository output: ", " | ", orders.MarshalString())

	if err != nil {
		l.Errorf("GetOrdersRepository error: ", " | ", err)
		return nil, err
	}

	for i := range orders {

		client, err := app.GetClientByID(orders[i].ClientID)
		l.Infof("GetClientByID output: ", " | ", client.MarshalString())

		if err != nil {
			l.Errorf("GetClientByID error: ", " | ", err)
			return nil, err
		}

		if err := app.GetAllOrderProductByIdService(orders[i].ID); err != nil {
			l.Errorf("GetAllOrderProductByIdService error: ", " | ", err)
			return nil, err
		}

		orderProductList, err := app.GetAllOrderProductByIdRepository(orders[i].ID)
		//l.Infof("GetAllOrderProductByIdRepository output: ", " | ", orderProductList.MarshalString())

		if err != nil {
			l.Errorf("GetAllOrderProductByIdRepository error: ", " | ", err)
			return nil, err
		}

		totalPrice := 0.0

		productList := make([]entity.ProductItem, 0)

		var voucher = entity.OutputVoucher{}

		if orders[i].VoucherID != nil {

			v, err := app.GetVoucherByID(*orders[i].VoucherID)
			l.Infof("GetVoucherByID output: ", " | ", v.MarshalString())

			if err != nil {
				l.Errorf("GetVoucherByID error: ", " | ", err)
				return nil, err
			}

			if v == nil {
				l.Infof("GetVoucherByID output: ", " | ", nil)
				return nil, errors.New("is not possible to save order because this voucher does not exist")
			}

			voucher = *v
		}

		for _, op := range orderProductList {
			if op.ProductID != nil {
				p, errGetC := app.GetProductByIDRepository(*op.ProductID)
				l.Infof("GetProductByIDRepository output: ", " | ", p.MarshalString())

				if errGetC != nil {
					l.Errorf("GetProductByIDRepository error: ", " | ", errGetC)
					return nil, errGetC
				}

				if p != nil {

					totalPrice = totalPrice + getTotalPrice(op.Quantity, p.Price)

					pp := entity.ProductItem{
						ID:            p.ID,
						Name:          p.Name,
						Quantity:      op.Quantity,
						Category:      p.Category.Value,
						Image:         p.Image,
						Description:   p.Description,
						Price:         p.Price,
						CreatedAt:     p.CreatedAt.Format(),
						DeactivatedAt: p.DeactivatedAt.Format(),
					}

					productList = append(productList, pp)
				}
			}
		}

		if voucher.Porcentage > 0 {
			totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
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
			Products:         productList,
			TotalPrice:       totalPrice,
			VoucherID:        orders[i].VoucherID,
			Status:           orders[i].Status.Value,
			VerificationCode: orders[i].VerificationCode.Value,
			CreatedAt:        orders[i].CreatedAt.Format(),
		}

		if strings.ToLower(order.Status) != valueObject.FinishedStatusKey {
			l.Infof("GetOrders output: ", " | ", order.MarshalString())
			orderList = append(orderList, order)
		}
	}
	l.Infof("GetOrders output: ", " | ", orderList)
	return orderList, nil
}

func (app hermesFoodsApp) GetOrderByID(id int64) (*entity.OutputOrder, error) {
	l.Infof("GetOrderByID: ", " | ", id)
	if err := app.orderService.GetOrderByID(id); err != nil {
		l.Errorf("GetOrderByIDService error: ", " | ", err)
		return nil, err
	}

	o, err := app.GetOrderByIDRepository(id)
	l.Infof("GetOrderByIDRepository output: ", " | ", o.MarshalString())

	if err != nil {
		l.Errorf("GetOrderByIDRepository error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		l.Infof("GetOrderByIDRepository output: ", " | ", nil)
		return nil, nil
	}

	outClient, err := app.GetClientByID(o.ClientID)
	l.Infof("GetClientByID output: ", " | ", outClient.MarshalString())

	if err != nil {
		l.Errorf("GetClientByID error: ", " | ", err)
		return nil, err
	}

	if err := app.GetAllOrderProductByIdService(id); err != nil {
		l.Errorf("GetAllOrderProductByIdService error: ", " | ", err)
		return nil, err
	}
	orderProductList, err := app.GetAllOrderProductByIdRepository(id)
	//l.Infof("GetAllOrderProductByIdRepository output: ", " | ", orderProductList.MarshalString())

	if err != nil {
		l.Errorf("GetAllOrderProductByIdRepository error: ", " | ", err)
		return nil, err
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = entity.OutputVoucher{}

	if o.VoucherID != nil {

		v, err := app.GetVoucherByID(*o.VoucherID)
		l.Infof("GetVoucherByID output: ", " | ", v.MarshalString())

		if err != nil {
			l.Errorf("GetVoucherByID error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			l.Infof("GetVoucherByID output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	for _, op := range orderProductList {
		if op.ProductID != nil {
			p, errGetC := app.GetProductByIDRepository(*op.ProductID)

			if errGetC != nil {
				l.Errorf("GetProductByIDRepository error: ", " | ", errGetC)
				return nil, errGetC
			}

			if p != nil {

				totalPrice = totalPrice + getTotalPrice(op.Quantity, p.Price)

				pp := entity.ProductItem{
					ID:            p.ID,
					Name:          p.Name,
					Quantity:      op.Quantity,
					Category:      p.Category.Value,
					Image:         p.Image,
					Description:   p.Description,
					Price:         p.Price,
					CreatedAt:     p.CreatedAt.Format(),
					DeactivatedAt: p.DeactivatedAt.Format(),
				}
				productList = append(productList, pp)
			}
		}
	}

	if outClient == nil {
		l.Infof("GetClientByID output: ", " | ", nil)
		return nil, errors.New("client is null")
	}

	if voucher.Porcentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
	}

	out := &entity.OutputOrder{
		ID:               o.ID,
		Client:           *outClient,
		Products:         productList,
		VoucherID:        o.VoucherID,
		TotalPrice:       totalPrice,
		Status:           o.Status.Value,
		VerificationCode: o.VerificationCode.Value,
		CreatedAt:        o.CreatedAt.Format(),
	}

	l.Infof("GetOrderByIDRepository output: ", " | ", out.MarshalString())
	return out, nil
}

func (app hermesFoodsApp) SaveOrder(order entity.Order) (*entity.OutputOrder, error) {
	l.Infof("SaveOrder: ", " | ", order.MarshalString())
	if err := app.GetClientByIDService(order.ClientID); err != nil {
		l.Errorf("GetClientByIDService error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByIDRepository(order.ClientID)
	l.Infof("GetClientByIDRepository output: ", " | ", c.MarshalString())

	if err != nil {
		l.Errorf("GetClientByIDRepository error: ", " | ", err)
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
	l.Infof("DoPaymentAPI output: ", " | ", out.MarshalString())

	if err != nil {
		l.Errorf("DoPaymentAPI error: ", " | ", err)
		return nil, err
	}

	if out.Error != nil {
		l.Errorf("DoPaymentAPI error: ", " | ", out.Error.Message, " | ", out.Error.Code)
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	log.Printf("output mercado pago api: %s\n", out.MarshalString())

	order.Status.Value = out.PaymentStatus

	o, err := app.SaveOrderService(order)

	if err != nil {
		l.Errorf("SaveOrderService error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		l.Infof("SaveOrderService output: ", " | ", nil)
		return nil, errors.New("is not possible to save order because it's null")
	}

	oRepo, err := app.SaveOrderRepository(*o)
	l.Infof("SaveOrderRepository output: ", " | ", oRepo.MarshalString())

	if err != nil {
		l.Errorf("SaveOrderRepository error: ", " | ", err)
		return nil, err
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = entity.OutputVoucher{}

	if order.VoucherID != nil {

		v, err := app.GetVoucherByID(*order.VoucherID)
		l.Infof("GetVoucherByID output: ", " | ", v.MarshalString())

		if err != nil {
			l.Errorf("GetVoucherByID error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			l.Infof("GetVoucherByID output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	for _, orderItems := range order.Items {
		if err := app.GetProductByIDService(orderItems.ProductID); err != nil {
			l.Errorf("GetProductByIDService error: ", " | ", err)
			return nil, err
		}

		product, err := app.GetProductByIDRepository(orderItems.ProductID)
		l.Infof("GetProductByIDRepository output: ", " | ", product.MarshalString())

		if err != nil {
			l.Errorf("GetProductByIDRepository error: ", " | ", err)
			return nil, err
		}

		if product == nil {
			l.Infof("GetProductByIDRepository output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because this product does not exists null")
		}

		pi := entity.ProductItem{
			ID:            product.ID,
			Name:          product.Name,
			Quantity:      orderItems.Quantity,
			Category:      product.Category.Value,
			Image:         product.Image,
			Description:   product.Description,
			Price:         product.Price,
			CreatedAt:     product.CreatedAt.Format(),
			DeactivatedAt: product.DeactivatedAt.Format(),
		}

		productList = append(productList, pi)

		tPrice := getTotalPrice(orderItems.Quantity, product.Price)

		totalPrice = totalPrice + tPrice

		var discount float64

		if voucher.Porcentage > 0 {
			discount = calculateDiscountByPercentage(voucher.Porcentage, tPrice)
		}

		opIn := entity.OrderProduct{
			Quantity:   orderItems.Quantity,
			TotalPrice: tPrice,
			OrderID:    oRepo.ID,
			ProductID:  &orderItems.ProductID,
			Discount:   discount,
		}

		opService, err := app.SaveOrderProductService(opIn)
		l.Infof("SaveOrderProductService output: ", " | ", opService.MarshalString())

		if err != nil {
			l.Errorf("SaveOrderProductService error: ", " | ", err)
			return nil, err
		}

		if opService == nil {
			l.Infof("SaveOrderProductService output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because it's null")
		}

		opRepo, err := app.SaveOrderProductRepository(opIn)
		l.Infof("SaveOrderProductRepository output: ", " | ", opRepo.MarshalString())

		if err != nil {
			l.Errorf("SaveOrderProductRepository error: ", " | ", err)
			return nil, err
		}

		if opRepo == nil {
			l.Infof("SaveOrderProductRepository output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because it's null")
		}

	}

	if voucher.Porcentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
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
		TotalPrice:       totalPrice,
		Products:         productList,
		VoucherID:        oRepo.VoucherID,
		Status:           oRepo.Status.Value,
		VerificationCode: oRepo.VerificationCode.Value,
		CreatedAt:        oRepo.CreatedAt.Format(),
	}

	l.Infof("SaveOrderRepository output: ", " | ", outOrder.MarshalString())
	return outOrder, nil
}

// ========== Product ==========

func (app hermesFoodsApp) SaveProduct(product entity.Product) (*entity.OutputProduct, error) {
	l.Infof("SaveProduct: ", " | ", product.MarshalString())
	p, err := app.SaveProductService(product)

	if err != nil {
		l.Errorf("SaveProductService error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		l.Infof("SaveProductService output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	pRepo, err := app.SaveProductRepository(*p)
	l.Infof("SaveProductRepository output: ", " | ", pRepo.MarshalString())

	if err != nil {
		l.Errorf("SaveProductRepository error: ", " | ", err)
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

	l.Infof("SaveProductRepository output: ", " | ", out.MarshalString())
	return out, nil
}

func (app hermesFoodsApp) GetProductByCategory(category string) ([]entity.OutputProduct, error) {
	l.Infof("GetProductByCategory: ", " | ", category)
	productList := make([]entity.OutputProduct, 0)

	if err := app.productService.GetProductByCategory(category); err != nil {
		l.Errorf("GetProductByCategoryService error: ", " | ", err)
		return nil, err
	}

	products, err := app.GetProductByCategoryRepository(category)
	//l.Infof("GetProductByCategoryRepository output: ", " | ", products.MarshalString())

	if err != nil {
		l.Errorf("GetProductByCategoryRepository error: ", " | ", err)
		return nil, err
	}

	if products == nil {
		l.Infof("GetProductByCategoryRepository output: ", " | ", nil)
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

	l.Infof("GetProductByCategoryRepository output: ", " | ", productList)
	return productList, nil
}

func (app hermesFoodsApp) UpdateProductByID(id int64, product entity.Product) (*entity.OutputProduct, error) {
	l.Infof("UpdateProductByID: ", " | ", id, " | ", product.MarshalString())

	if err := app.GetProductByIDService(id); err != nil {
		l.Errorf("GetProductByIDService error: ", " | ", err)
		return nil, err
	}

	pByID, err := app.GetProductByIDRepository(id)
	l.Infof("GetProductByIDRepository output: ", " | ", pByID.MarshalString())

	if err != nil {
		l.Errorf("GetProductByIDRepository error: ", " | ", err)
		return nil, err
	}

	if pByID == nil {
		l.Infof("GetProductByIDRepository output: ", " | ", nil)
		return nil, errors.New("was not found any product with this id")
	}

	p, err := app.UpdateProductByIDService(id, product)
	l.Infof("UpdateProductByIDService output: ", " | ", p.MarshalString())

	if err != nil {
		l.Errorf("UpdateProductByIDService error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		l.Infof("UpdateProductByIDService output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	pRepo, err := app.UpdateProductByIDRepository(id, *p)
	l.Infof("UpdateProductByIDRepository output: ", " | ", pRepo.MarshalString())

	if err != nil {
		l.Errorf("UpdateProductByIDRepository error: ", " | ", err)
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

	l.Infof("UpdateProductByIDRepository output: ", " | ", out.MarshalString())
	return out, nil
}

func (app hermesFoodsApp) DeleteProductByID(id int64) error {
	l.Infof("DeleteProductByID: ", " | ", id)

	if err := app.GetProductByIDService(id); err != nil {
		l.Errorf("GetProductByIDService error: ", " | ", err)
		return err
	}

	pByID, err := app.GetProductByIDRepository(id)
	l.Infof("GetProductByIDRepository output: ", " | ", pByID.MarshalString())

	if err != nil {
		l.Errorf("GetProductByIDRepository error: ", " | ", err)
		return err
	}

	if pByID == nil {
		l.Infof("GetProductByIDRepository output: ", " | ", nil)
		return errors.New("was not found any product with this id")
	}

	if err := app.DeleteProductByIDService(id); err != nil {
		l.Errorf("DeleteProductByIDService error: ", " | ", err)
		return err
	}

	l.Infof("DeleteProductByIDRepository output: ", " | ", nil)
	return app.DeleteProductByIDRepository(id)

}

// ========== Voucher ==========

func (app hermesFoodsApp) SaveVoucher(voucher entity.Voucher) (*entity.OutputVoucher, error) {
	l.Infof("SaveVoucher: ", " | ", voucher.MarshalString())
	voucherSvc, err := app.SaveVoucherService(voucher)

	if err != nil {
		l.Errorf("SaveVoucherService error: ", " | ", err)
		return nil, err
	}

	if voucherSvc == nil {
		l.Infof("SaveVoucherService output: ", " | ", nil)
		return nil, errors.New("is not possible to save voucher because it's null")
	}

	rVoucher, err := app.SaveVoucherRepository(voucher)
	l.Infof("SaveVoucherRepository output: ", " | ", rVoucher.MarshalString())

	if err != nil {
		l.Errorf("SaveVoucherRepository error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		l.Infof("SaveVoucherRepository output: ", " | ", nil)
		return nil, errors.New("was not possible to save voucher because it's null")
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("SaveVoucherRepository output: ", " | ", vOut.MarshalString())
	return &vOut, nil
}

func (app hermesFoodsApp) GetVoucherByID(id int64) (*entity.OutputVoucher, error) {
	l.Infof("GetVoucherByID: ", " | ", id)
	if err := app.GetVoucherByIDService(id); err != nil {
		l.Errorf("GetVoucherByIDService error: ", " | ", err)
		return nil, err
	}

	rVoucher, err := app.GetVoucherByIDRepository(id)
	l.Infof("GetVoucherByIDRepository output: ", " | ", rVoucher.MarshalString())

	if err != nil {
		l.Errorf("GetVoucherByIDRepository error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		l.Infof("GetVoucherByIDRepository output: ", " | ", nil)
		return nil, fmt.Errorf("voucher not found with the %d id", id)
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("GetVoucherByIDRepository output: ", " | ", vOut.MarshalString())
	return &vOut, nil

}

func (app hermesFoodsApp) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.OutputVoucher, error) {
	l.Infof("UpdateVoucherByID: ", " | ", id, " | ", voucher.MarshalString())
	voucherSvc, err := app.UpdateVoucherByIDService(id, voucher)

	if err != nil {
		l.Errorf("UpdateVoucherByIDService error: ", " | ", err)
		return nil, err
	}

	if voucherSvc == nil {
		l.Infof("UpdateVoucherByIDService output: ", " | ", nil)
		return nil, errors.New("is not possible to update voucher because it's null")
	}

	rVoucher, err := app.UpdateVoucherByIDRepository(id, voucher)
	l.Infof("UpdateVoucherByIDRepository output: ", " | ", rVoucher.MarshalString())

	if err != nil {
		l.Errorf("UpdateVoucherByIDRepository error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		l.Infof("UpdateVoucherByIDRepository output: ", " | ", nil)
		return nil, errors.New("was not possible to update voucher because it's null")
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("UpdateVoucherByIDRepository output: ", " | ", vOut.MarshalString())
	return &vOut, nil
}

func getTotalPrice(quantity int64, productPrice float64) float64 {
	l.Infof("getTotalPrice: ", " | ", quantity, " | ", productPrice)
	return productPrice * float64(quantity)
}

func calculateDiscountByPercentage(percentage int64, value float64) float64 {
	if percentage == 0 {
		return value
	}

	l.Infof("calculateDiscountByPercentage: ", " | ", percentage, " | ", value)
	return value - (value * (float64(percentage) / 100))
}

func (app hermesFoodsApp) DoPaymentAPI(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error) {
	l.Infof("DoPaymentAPI: ", " | ", input.MarshalString())
	return app.paymentAPI.DoPayment(ctx, input)
}
