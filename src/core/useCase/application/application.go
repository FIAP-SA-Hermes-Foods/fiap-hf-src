package application

import (
	"context"
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
	l "fiap-hf-src/src/external/logger"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"fmt"
	"strings"
)

type HermesFoodsApp interface {

	// Client Methods
	SaveClient(client entity.Client) (*dto.OutputClient, error)
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	GetClientByID(id int64) (*dto.OutputClient, error)

	// Order Methods
	SaveOrder(order entity.Order) (*dto.OutputOrder, error)
	GetOrderByID(id int64) (*dto.OutputOrder, error)
	UpdateOrderByID(id int64, order entity.Order) (*dto.OutputOrder, error)
	GetOrders() ([]dto.OutputOrder, error)

	// Product Methods
	SaveProduct(product entity.Product) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	UpdateProductByID(id int64, product entity.Product) (*dto.OutputProduct, error)
	DeleteProductByID(id int64) error

	// Voucher Methods
	SaveVoucher(voucher entity.Voucher) (*dto.OutputVoucher, error)
	GetVoucherByID(id int64) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*dto.OutputVoucher, error)
}

type hermesFoodsApp struct {
	Ctx                 context.Context
	paymentAPI          interfaces.PaymentAPI
	clientRepo          interfaces.ClientRepository
	clientService       interfaces.ClientService
	orderRepo           interfaces.OrderRepository
	orderService        interfaces.OrderService
	orderProductRepo    interfaces.OrderProductRepository
	orderProductService interfaces.OrderProductService
	productRepo         interfaces.ProductRepository
	productService      interfaces.ProductService
	voucherRepo         interfaces.VoucherRepository
	voucherService      interfaces.VoucherService
}

func NewHermesFoodsApp(ctx context.Context, paymentAPI interfaces.PaymentAPI, clientRepo interfaces.ClientRepository, orderRepo interfaces.OrderRepository, orderProductRepo interfaces.OrderProductRepository, productRepo interfaces.ProductRepository, voucherRepo interfaces.VoucherRepository, clientService interfaces.ClientService, orderService interfaces.OrderService, orderProductService interfaces.OrderProductService, productService interfaces.ProductService, voucherService interfaces.VoucherService) HermesFoodsApp {
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

func (app hermesFoodsApp) GetClientByID(id int64) (*dto.OutputClient, error) {
	l.Infof("GetClientByID: ", " | ", id)
	if err := app.GetClientByIDService(id); err != nil {
		l.Errorf("GetClientByID error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByIDRepository(id)

	if err != nil {
		l.Errorf("GetClientByID error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByID output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	l.Infof("GetClientByID output: ", " | ", ps.MarshalString(out))
	return out, err
}

func (app hermesFoodsApp) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	l.Infof("GetClientByCPF: ", " | ", cpf)
	if err := app.GetClientByCPFService(cpf); err != nil {
		l.Errorf("GetClientByCPF error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByCPFRepository(cpf)

	if err != nil {
		l.Errorf("GetClientByCPF error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByCPF output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	l.Infof("GetClientByCPF output: ", " | ", ps.MarshalString(out))
	return out, err
}

func (app hermesFoodsApp) SaveClient(client entity.Client) (*dto.OutputClient, error) {
	l.Infof("SaveClient: ", " | ", ps.MarshalString(client))
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

	if err != nil {
		l.Errorf("SaveClient error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("SaveClient output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because it's null")
	}

	cRepo, err := app.SaveClientRepository(*c)
	l.Infof("SaveClient output: ", " | ", ps.MarshalString(cRepo))

	if err != nil {
		l.Errorf("SaveClient error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputClient{
		ID:        cRepo.ID,
		Name:      cRepo.Name,
		CPF:       cRepo.CPF.Value,
		Email:     cRepo.Email,
		CreatedAt: cRepo.CreatedAt.Format(),
	}

	l.Infof("SaveClient output: ", " | ", ps.MarshalString(out))
	return out, nil
}

// ========== Order ==========

func (app hermesFoodsApp) UpdateOrderByID(id int64, order entity.Order) (*dto.OutputOrder, error) {
	l.Infof("UpdateOrderByID: ", " | ", id, " | ", ps.MarshalString(order))
	oSvc, err := app.UpdateOrderByIDService(id, order)

	if err != nil {
		l.Errorf("UpdateOrderByID error: ", " | ", err)
		return nil, err
	}

	if oSvc == nil {
		l.Infof("UpdateOrderByID output: ", " | ", nil)
		return nil, errors.New("order is null, is not possible to proceed with update order")
	}

	oRepo, err := app.UpdateOrderByIDRepository(id, order)

	if err != nil {
		l.Errorf("UpdateOrderByID error: ", " | ", err)
		return nil, err
	}

	client, err := app.GetClientByID(oRepo.ClientID)

	if err != nil {
		l.Errorf("UpdateOrderByID error: ", " | ", err)
		return nil, err
	}

	if client == nil {
		l.Infof("UpdateOrderByID output: ", " | ", nil)
		return nil, errors.New("client is null, is not possible to proceed with update order")
	}

	out := &dto.OutputOrder{
		ID:               oRepo.ID,
		Client:           *client,
		VoucherID:        oRepo.VoucherID,
		Status:           oRepo.Status.Value,
		VerificationCode: oRepo.VerificationCode.Value,
		CreatedAt:        oRepo.CreatedAt.Format(),
	}

	l.Infof("UpdateOrderByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) GetOrders() ([]dto.OutputOrder, error) {
	l.Infof("GetOrders: ", " | ")
	orderList := make([]dto.OutputOrder, 0)

	orders, err := app.GetOrdersRepository()

	if err != nil {
		l.Errorf("GetOrders error: ", " | ", err)
		return nil, err
	}

	for i := range orders {

		client, err := app.GetClientByID(orders[i].ClientID)

		if err != nil {
			l.Errorf("GetOrders error: ", " | ", err)
			return nil, err
		}

		if err := app.GetAllOrderProductByIdService(orders[i].ID); err != nil {
			l.Errorf("GetOrders error: ", " | ", err)
			return nil, err
		}

		orderProductList, err := app.GetAllOrderProductByIdRepository(orders[i].ID)

		if err != nil {
			l.Errorf("GetOrders error: ", " | ", err)
			return nil, err
		}

		totalPrice := 0.0

		productList := make([]entity.ProductItem, 0)

		var voucher = dto.OutputVoucher{}

		if orders[i].VoucherID != nil {

			v, err := app.GetVoucherByID(*orders[i].VoucherID)

			if err != nil {
				l.Errorf("GetOrders error: ", " | ", err)
				return nil, err
			}

			if v == nil {
				l.Infof("GetOrders output: ", " | ", nil)
				return nil, errors.New("is not possible to save order because this voucher does not exist")
			}

			voucher = *v
		}

		for _, op := range orderProductList {
			if op.ProductID != nil {
				p, errGetC := app.GetProductByIDRepository(*op.ProductID)

				if errGetC != nil {
					l.Errorf("GetOrders error: ", " | ", errGetC)
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

		order := dto.OutputOrder{
			ID: orders[i].ID,
			Client: dto.OutputClient{
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

		if strings.ToLower(order.Status) != com.FinishedStatusKey {
			orderList = append(orderList, order)
		}
	}

	l.Infof("GetOrders output: ", " | ", orderList)
	return orderList, nil
}

func (app hermesFoodsApp) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	l.Infof("GetOrderByID: ", " | ", id)
	if err := app.orderService.GetOrderByID(id); err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	o, err := app.GetOrderByIDRepository(id)

	if err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		l.Infof("GetOrderByID output: ", " | ", nil)
		return nil, nil
	}

	outClient, err := app.GetClientByID(o.ClientID)

	if err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	if err := app.GetAllOrderProductByIdService(id); err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}
	orderProductList, err := app.GetAllOrderProductByIdRepository(id)

	if err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = dto.OutputVoucher{}

	if o.VoucherID != nil {

		v, err := app.GetVoucherByID(*o.VoucherID)

		if err != nil {
			l.Errorf("GetOrderByID error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			voucherNullErr := errors.New("is not possible to save order because this voucher does not exist")
			l.Errorf("GetOrderByID error: ", " | ", voucherNullErr)
			return nil, voucherNullErr
		}

		voucher = *v
	}

	for _, op := range orderProductList {
		if op.ProductID != nil {
			p, errGetC := app.GetProductByIDRepository(*op.ProductID)

			if errGetC != nil {
				l.Errorf("GetOrderByID error: ", " | ", errGetC)
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
		clientNullErr := errors.New("client is null")
		l.Errorf("GetOrderByID error: ", " | ", clientNullErr)
		return nil, clientNullErr
	}

	if voucher.Porcentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
	}

	out := &dto.OutputOrder{
		ID:               o.ID,
		Client:           *outClient,
		Products:         productList,
		VoucherID:        o.VoucherID,
		TotalPrice:       totalPrice,
		Status:           o.Status.Value,
		VerificationCode: o.VerificationCode.Value,
		CreatedAt:        o.CreatedAt.Format(),
	}

	l.Infof("GetOrderByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) SaveOrder(order entity.Order) (*dto.OutputOrder, error) {
	l.Infof("SaveOrder: ", " | ", ps.MarshalString(order))
	if err := app.GetClientByIDService(order.ClientID); err != nil {
		l.Errorf("SaveOrder error: ", " | ", err)
		return nil, err
	}

	c, err := app.GetClientByIDRepository(order.ClientID)

	if err != nil {
		l.Errorf("SaveOrder error: ", " | ", err)
		return nil, err
	}

	inputDoPaymentAPI := dto.InputPaymentAPI{
		Price: 0.0,
		Client: entity.Client{
			ID:   c.ID,
			Name: c.Name,
			CPF: com.Cpf{
				Value: c.CPF.Value,
			},
			Email: c.Email,
			CreatedAt: com.CreatedAt{
				Value: c.CreatedAt.Value,
			},
		},
	}

	out, err := app.DoPaymentAPI(app.Ctx, inputDoPaymentAPI)

	if err != nil {
		l.Errorf("SaveOrder error: ", " | ", err)
		return nil, err
	}

	if out.Error != nil {
		l.Errorf("SaveOrder error: ", " | ", out.Error.Message, " | ", out.Error.Code)
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	order.Status.Value = out.PaymentStatus

	o, err := app.SaveOrderService(order)

	if err != nil {
		l.Errorf("SaveOrder error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		orderNullErr := errors.New("is not possible to save order because it's null")
		l.Infof("SaveOrder output: ", " | ", orderNullErr)
		return nil, orderNullErr
	}

	oRepo, err := app.SaveOrderRepository(*o)
	l.Infof("SaveOrder output: ", " | ", ps.MarshalString(oRepo))

	if err != nil {
		l.Errorf("SaveOrder error: ", " | ", err)
		return nil, err
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = dto.OutputVoucher{}

	if order.VoucherID != nil {

		v, err := app.GetVoucherByID(*order.VoucherID)
		l.Infof("SaveOrder output: ", " | ", ps.MarshalString(v))

		if err != nil {
			l.Errorf("SaveOrder error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			voucherNullErr := errors.New("is not possible to save order because this voucher does not exist")
			l.Infof("SaveOrder output: ", " | ", voucherNullErr)
			return nil, voucherNullErr
		}

		voucher = *v
	}

	for _, orderItems := range order.Items {
		if err := app.GetProductByIDService(orderItems.ProductID); err != nil {
			l.Errorf("SaveOrder error: ", " | ", err)
			return nil, err
		}

		product, err := app.GetProductByIDRepository(orderItems.ProductID)

		if err != nil {
			l.Errorf("SaveOrder error: ", " | ", err)
			return nil, err
		}

		if product == nil {
			productNullErr := errors.New("is not possible to save order because this product does not exists null")
			l.Infof("SaveOrder output: ", " | ", productNullErr)
			return nil, productNullErr
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
			l.Errorf("SaveOrder error: ", " | ", err)
			return nil, err
		}

		if opService == nil {
			orderProductNullErr := errors.New("is not possible to save order because it's null")
			l.Infof("SaveOrder output: ", " | ", orderProductNullErr)
			return nil, orderProductNullErr
		}

		opRepo, err := app.SaveOrderProductRepository(opIn)

		if err != nil {
			l.Errorf("SaveOrder error: ", " | ", err)
			return nil, err
		}

		if opRepo == nil {
			l.Infof("SaveOrder output: ", " | ", nil)
			return nil, errors.New("is not possible to save order because it's null")
		}

	}

	if voucher.Porcentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
	}

	outClient := dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	outOrder := &dto.OutputOrder{
		ID:               oRepo.ID,
		Client:           outClient,
		TotalPrice:       totalPrice,
		Products:         productList,
		VoucherID:        oRepo.VoucherID,
		Status:           oRepo.Status.Value,
		VerificationCode: oRepo.VerificationCode.Value,
		CreatedAt:        oRepo.CreatedAt.Format(),
	}

	l.Infof("SaveOrder output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

// ========== Product ==========

func (app hermesFoodsApp) SaveProduct(product entity.Product) (*dto.OutputProduct, error) {
	l.Infof("SaveProduct: ", " | ", ps.MarshalString(product))
	p, err := app.SaveProductService(product)

	if err != nil {
		l.Errorf("SaveProduct error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		l.Infof("SaveProduct output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	pRepo, err := app.SaveProductRepository(*p)

	if err != nil {
		l.Errorf("SaveProduct error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputProduct{
		ID:            pRepo.ID,
		Name:          pRepo.Name,
		Category:      pRepo.Category.Value,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt.Format(),
		DeactivatedAt: pRepo.DeactivatedAt.Format(),
	}

	l.Infof("SaveProduct output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	l.Infof("GetProductByCategory: ", " | ", category)
	productList := make([]dto.OutputProduct, 0)

	if err := app.productService.GetProductByCategory(category); err != nil {
		l.Errorf("GetProductByCategory error: ", " | ", err)
		return nil, err
	}

	products, err := app.GetProductByCategoryRepository(category)

	if err != nil {
		l.Errorf("GetProductByCategory error: ", " | ", err)
		return nil, err
	}

	if products == nil {
		l.Infof("GetProductByCategory output: ", " | ", nil)
		return nil, nil
	}

	for i := range products {
		product := dto.OutputProduct{
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

	l.Infof("GetProductByCategory output: ", " | ", productList)
	return productList, nil
}

func (app hermesFoodsApp) UpdateProductByID(id int64, product entity.Product) (*dto.OutputProduct, error) {
	l.Infof("UpdateProductByID: ", " | ", id, " | ", ps.MarshalString(product))

	if err := app.GetProductByIDService(id); err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	pByID, err := app.GetProductByIDRepository(id)

	if err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	if pByID == nil {
		productNullErr := errors.New("was not found any product with this id")
		l.Infof("UpdateProductByID output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	p, err := app.UpdateProductByIDService(id, product)

	if err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		productNullErr := errors.New("is not possible to save product because it's null")
		l.Errorf("UpdateProductByID output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	pRepo, err := app.UpdateProductByIDRepository(id, *p)

	if err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputProduct{
		ID:            pRepo.ID,
		Name:          pRepo.Name,
		Category:      pRepo.Category.Value,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt.Format(),
		DeactivatedAt: pRepo.DeactivatedAt.Format(),
	}

	l.Infof("UpdateProductByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) DeleteProductByID(id int64) error {
	l.Infof("DeleteProductByID: ", " | ", id)

	if err := app.GetProductByIDService(id); err != nil {
		l.Errorf("DeleteProductByID error: ", " | ", err)
		return err
	}

	pByID, err := app.GetProductByIDRepository(id)

	if err != nil {
		l.Errorf("DeleteProductByID error: ", " | ", err)
		return err
	}

	if pByID == nil {
		productNullErr := errors.New("was not found any product with this id")
		l.Infof("DeleteProductByID output: ", " | ", productNullErr)
		return productNullErr
	}

	if err := app.DeleteProductByIDService(id); err != nil {
		l.Errorf("DeleteProductByID error: ", " | ", err)
		return err
	}

	l.Infof("DeleteProductByID output: ", " | ", nil)
	return app.DeleteProductByIDRepository(id)
}

// ========== Voucher ==========

func (app hermesFoodsApp) SaveVoucher(voucher entity.Voucher) (*dto.OutputVoucher, error) {
	l.Infof("SaveVoucher: ", " | ", ps.MarshalString(voucher))
	voucherSvc, err := app.SaveVoucherService(voucher)

	if err != nil {
		l.Errorf("SaveVoucher error: ", " | ", err)
		return nil, err
	}

	if voucherSvc == nil {
		voucherNullErr := errors.New("is not possible to save voucher because it's null")
		l.Errorf("SaveVoucher output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	rVoucher, err := app.SaveVoucherRepository(voucher)

	if err != nil {
		l.Errorf("SaveVoucher error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNullErr := errors.New("is not possible to save voucher because it's null")
		l.Errorf("SaveVoucher output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("SaveVoucher output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app hermesFoodsApp) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {
	l.Infof("GetVoucherByID: ", " | ", id)
	if err := app.GetVoucherByIDService(id); err != nil {
		l.Errorf("GetVoucherByID error: ", " | ", err)
		return nil, err
	}

	rVoucher, err := app.GetVoucherByIDRepository(id)

	if err != nil {
		l.Errorf("GetVoucherByID error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNotFoundErr := fmt.Errorf("voucher not found with the %d id", id)
		l.Errorf("GetVoucherByID output: ", " | ", voucherNotFoundErr)
		return nil, voucherNotFoundErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("GetVoucherByID output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app hermesFoodsApp) UpdateVoucherByID(id int64, voucher entity.Voucher) (*dto.OutputVoucher, error) {
	l.Infof("UpdateVoucherByID: ", " | ", id, " | ", ps.MarshalString(voucher))
	voucherSvc, err := app.UpdateVoucherByIDService(id, voucher)

	if err != nil {
		l.Errorf("UpdateVoucherByID error: ", " | ", err)
		return nil, err
	}

	if voucherSvc == nil {
		voucherNullErr := errors.New("is not possible to update voucher because it's null")
		l.Infof("UpdateVoucherByID output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	rVoucher, err := app.UpdateVoucherByIDRepository(id, voucher)

	if err != nil {
		l.Errorf("UpdateVoucherByID error: ", " | ", err)
		return nil, err
	}

	if rVoucher == nil {
		voucherNullErr := errors.New("is not possible to update voucher because it's null")
		l.Infof("UpdateVoucherByID output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	vOut := dto.OutputVoucher{
		ID:         rVoucher.ID,
		Code:       rVoucher.Code,
		Porcentage: rVoucher.Porcentage,
		CreatedAt:  rVoucher.CreatedAt.Format(),
		ExpiresAt:  rVoucher.ExpiresAt.Format(),
	}

	l.Infof("UpdateVoucherByID output: ", " | ", ps.MarshalString(vOut))
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

func (app hermesFoodsApp) DoPaymentAPI(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error) {
	return app.paymentAPI.DoPayment(ctx, input)
}
