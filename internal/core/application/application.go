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

// ========== Order ==========

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

		if err := app.GetAllOrderProductByIdService(orders[i].ID); err != nil {
			return nil, err
		}
		orderProductList, err := app.GetAllOrderProductByIdRepository(orders[i].ID)

		if err != nil {
			return nil, err
		}

		totalPrice := 0.0

		productList := make([]entity.ProductItem, 0)

		var voucher = entity.OutputVoucher{}

		if orders[i].VoucherID != nil {

			v, err := app.GetVoucherByID(*orders[i].VoucherID)

			if err != nil {
				return nil, err
			}

			if v == nil {
				return nil, errors.New("is not possible to save order because this voucher does not exist")
			}

			voucher = *v
		}

		for _, op := range orderProductList {
			if op.ProductID != nil {
				p, errGetC := app.GetProductByIDRepository(*op.ProductID)

				if errGetC != nil {
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

	if err := app.GetAllOrderProductByIdService(id); err != nil {
		return nil, err
	}
	orderProductList, err := app.GetAllOrderProductByIdRepository(id)

	if err != nil {
		return nil, err
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = entity.OutputVoucher{}

	if o.VoucherID != nil {

		v, err := app.GetVoucherByID(*o.VoucherID)

		if err != nil {
			return nil, err
		}

		if v == nil {
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	for _, op := range orderProductList {
		if op.ProductID != nil {
			p, errGetC := app.GetProductByIDRepository(*op.ProductID)

			if errGetC != nil {
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

	return out, nil
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

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = entity.OutputVoucher{}

	if order.VoucherID != nil {

		v, err := app.GetVoucherByID(*order.VoucherID)

		if err != nil {
			return nil, err
		}

		if v == nil {
			return nil, errors.New("is not possible to save order because this voucher does not exist")
		}

		voucher = *v
	}

	for _, orderItems := range order.Items {
		if err := app.GetProductByIDService(orderItems.ProductID); err != nil {
			return nil, err
		}

		product, err := app.GetProductByIDRepository(orderItems.ProductID)

		if err != nil {
			return nil, err
		}

		if product == nil {
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

		if err != nil {
			return nil, err
		}

		if opService == nil {
			return nil, errors.New("is not possible to save order because it's null")
		}

		opRepo, err := app.SaveOrderProductRepository(opIn)

		if err != nil {
			return nil, err
		}

		if opRepo == nil {
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

	return outOrder, nil
}

// ========== Product ==========

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

// ========== Voucher ==========

func (app hermesFoodsApp) SaveVoucher(voucher entity.Voucher) (*entity.OutputVoucher, error) {
	voucherSvc, err := app.SaveVoucherService(voucher)

	if err != nil {
		return nil, err
	}

	if voucherSvc == nil {
		return nil, errors.New("is not possible to save voucher because it's null")
	}

	rVoucher, err := app.SaveVoucherRepository(voucher)

	if err != nil {
		return nil, err
	}

	if rVoucher == nil {
		return nil, errors.New("was not possible to save voucher because it's null")
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	return &vOut, nil
}

func (app hermesFoodsApp) GetVoucherByID(id int64) (*entity.OutputVoucher, error) {
	if err := app.GetVoucherByIDService(id); err != nil {
		return nil, err
	}

	rVoucher, err := app.GetVoucherByIDRepository(id)

	if err != nil {
		return nil, err
	}

	if rVoucher == nil {
		return nil, fmt.Errorf("voucher not found with the %d id", id)
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	return &vOut, nil

}

func (app hermesFoodsApp) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.OutputVoucher, error) {
	voucherSvc, err := app.UpdateVoucherByIDService(id, voucher)

	if err != nil {
		return nil, err
	}

	if voucherSvc == nil {
		return nil, errors.New("is not possible to update voucher because it's null")
	}

	rVoucher, err := app.UpdateVoucherByIDRepository(id, voucher)

	if err != nil {
		return nil, err
	}

	if rVoucher == nil {
		return nil, errors.New("was not possible to update voucher because it's null")
	}

	vOut := entity.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	return &vOut, nil
}

func getTotalPrice(quantity int64, productPrice float64) float64 {
	return productPrice * float64(quantity)
}

func calculateDiscountByPercentage(percentage int64, value float64) float64 {
	if percentage == 0 {
		return value
	}

	return value - (value * (float64(percentage) / 100))
}

func (app hermesFoodsApp) DoPaymentAPI(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error) {
	return app.paymentAPI.DoPayment(ctx, input)
}
