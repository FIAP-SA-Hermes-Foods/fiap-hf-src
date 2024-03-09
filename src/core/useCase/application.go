package useCase

import (
	"context"
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	l "fiap-hf-src/src/base/logger"
	"fiap-hf-src/src/core/entity"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"fmt"
	"strings"
)

var _ interfaces.HermesFoodsUseCase = (*hermesFoodsApp)(nil)

type hermesFoodsApp struct {
	ctx                 context.Context
	paymentAPI          interfaces.PaymentUseCase
	clientUseCase       interfaces.ClientUseCase
	orderUseCase        interfaces.OrderUseCase
	orderProductUseCase interfaces.OrderProductUseCase
	productUseCase      interfaces.ProductUseCase
	voucherUseCase      interfaces.VoucherUseCase
}

func NewHermesFoodsApp(ctx context.Context, paymentAPI interfaces.PaymentUseCase, clientUseCase interfaces.ClientUseCase, orderUseCase interfaces.OrderUseCase, orderProductUseCase interfaces.OrderProductUseCase, productUseCase interfaces.ProductUseCase, voucherUseCase interfaces.VoucherUseCase) *hermesFoodsApp {
	return &hermesFoodsApp{
		ctx:                 ctx,
		paymentAPI:          paymentAPI,
		clientUseCase:       clientUseCase,
		orderUseCase:        orderUseCase,
		orderProductUseCase: orderProductUseCase,
		productUseCase:      productUseCase,
		voucherUseCase:      voucherUseCase,
	}
}

// ========== Client ==========

func (app hermesFoodsApp) GetClientByID(id int64) (*dto.OutputClient, error) {

	c, err := app.clientUseCase.GetClientByID(id)

	if err != nil {
		l.Errorf("GetClientByID error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByID output: ", " | ", nil)
		return nil, nil
	}

	l.Infof("GetClientByID output: ", " | ", ps.MarshalString(c))
	return c, err
}

func (app hermesFoodsApp) GetClientByCPF(cpf string) (*dto.OutputClient, error) {

	c, err := app.clientUseCase.GetClientByCPF(cpf)

	if err != nil {
		l.Errorf("GetClientByCPF error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("GetClientByCPF output: ", " | ", nil)
		return nil, nil
	}

	l.Infof("GetClientByCPF output: ", " | ", ps.MarshalString(c))
	return c, err
}

func (app hermesFoodsApp) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	l.Infof("SaveClient: ", " | ", ps.MarshalString(client))
	clientWithCpf, err := app.GetClientByCPF(client.CPF)

	if err != nil {
		l.Errorf("SaveClient error: ", " | ", err)
		return nil, err
	}

	if clientWithCpf != nil {
		l.Infof("SaveClient output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

	c, err := app.clientUseCase.SaveClient(client)

	if err != nil {
		l.Errorf("SaveClient error: ", " | ", err)
		return nil, err
	}

	if c == nil {
		l.Infof("SaveClient output: ", " | ", nil)
		return nil, errors.New("is not possible to save client because it's null")
	}

	l.Infof("SaveClient output: ", " | ", ps.MarshalString(c))
	return c, nil
}

// ========== Order ==========

func (app hermesFoodsApp) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error) {
	l.Infof("UpdateOrderByID: ", " | ", id, " | ", ps.MarshalString(order))
	oSvc, err := app.orderUseCase.UpdateOrderByID(id, order)

	if err != nil {
		l.Errorf("UpdateOrderByID error: ", " | ", err)
		return nil, err
	}

	if oSvc == nil {
		l.Infof("UpdateOrderByID output: ", " | ", nil)
		return nil, errors.New("order is null, is not possible to proceed with update order")
	}

	client, err := app.GetClientByID(oSvc.Client.ID)

	if err != nil {
		l.Errorf("UpdateOrderByID error: ", " | ", err)
		return nil, err
	}

	if client == nil {
		l.Infof("UpdateOrderByID output: ", " | ", nil)
		return nil, errors.New("client is null, is not possible to proceed with update order")
	}

	out := &dto.OutputOrder{
		ID:               oSvc.ID,
		Client:           *client,
		VoucherID:        oSvc.VoucherID,
		Status:           oSvc.Status,
		VerificationCode: oSvc.VerificationCode,
		CreatedAt:        oSvc.CreatedAt,
	}

	l.Infof("UpdateOrderByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) GetOrders() ([]dto.OutputOrder, error) {
	l.Infof("GetOrders: ", " | ")
	orderList := make([]dto.OutputOrder, 0)

	orders, err := app.orderUseCase.GetOrders()

	if err != nil {
		l.Errorf("GetOrders error: ", " | ", err)
		return nil, err
	}

	for i := range orders {

		client, err := app.GetClientByID(orders[i].Client.ID)

		if err != nil {
			l.Errorf("GetOrders error: ", " | ", err)
			return nil, err
		}

		orderProductList, err := app.orderProductUseCase.GetAllOrderProductByOrderID(orders[i].ID)

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
				p, errGetC := app.productUseCase.GetProductByID(*op.ProductID)

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
						Category:      p.Category,
						Image:         p.Image,
						Description:   p.Description,
						Price:         p.Price,
						CreatedAt:     p.CreatedAt,
						DeactivatedAt: p.DeactivatedAt,
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
			Status:           orders[i].Status,
			VerificationCode: orders[i].VerificationCode,
			CreatedAt:        orders[i].CreatedAt,
		}

		if strings.ToLower(order.Status) != com.FinishedStatusKey {
			orderList = append(orderList, order)
		}
	}

	l.Infof("GetOrders output: ", " | ", ps.MarshalString(orderList))
	return orderList, nil
}

func (app hermesFoodsApp) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	l.Infof("GetOrderByID: ", " | ", id)

	o, err := app.orderUseCase.GetOrderByID(id)

	if err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		l.Infof("GetOrderByID output: ", " | ", nil)
		return nil, nil
	}

	outClient, err := app.GetClientByID(o.Client.ID)

	if err != nil {
		l.Errorf("GetOrderByID error: ", " | ", err)
		return nil, err
	}

	orderProductList, err := app.orderProductUseCase.GetAllOrderProductByOrderID(id)

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
			p, errGetC := app.productUseCase.GetProductByID(*op.ProductID)

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
					Category:      p.Category,
					Image:         p.Image,
					Description:   p.Description,
					Price:         p.Price,
					CreatedAt:     p.CreatedAt,
					DeactivatedAt: p.DeactivatedAt,
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
		Status:           o.Status,
		VerificationCode: o.VerificationCode,
		CreatedAt:        o.CreatedAt,
	}

	l.Infof("GetOrderByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	l.Infof("SaveOrderApp: ", " | ", ps.MarshalString(order))

	c, err := app.clientUseCase.GetClientByID(order.ClientID)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	inputDoPaymentAPI := dto.InputPaymentAPI{
		// Price: 0.0,
		Client: entity.Client{
			ID:   c.ID,
			Name: c.Name,
			CPF: com.Cpf{
				Value: c.CPF,
			},
			Email:     c.Email,
			CreatedAt: com.CreatedAt{
				// Value: c.CreatedAt,
			},
		},
	}

	out, err := app.DoPaymentAPI(app.ctx, inputDoPaymentAPI)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	if out.Error != nil {
		l.Errorf("SaveOrderApp error: ", " | ", out.Error.Message, " | ", out.Error.Code)
		return nil, fmt.Errorf("error to do payment message: %s, code: %s", out.Error.Message, out.Error.Code)
	}

	order.Status = out.PaymentStatus

	o, err := app.orderUseCase.SaveOrder(order)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		orderNullErr := errors.New("is not possible to save order because it's null")
		l.Infof("SaveOrderApp output: ", " | ", orderNullErr)
		return nil, orderNullErr
	}

	productList := make([]entity.ProductItem, 0)

	totalPrice := 0.0

	var voucher = dto.OutputVoucher{}

	if order.VoucherID != nil {

		v, err := app.GetVoucherByID(*order.VoucherID)
		l.Infof("SaveOrderApp output: ", " | ", ps.MarshalString(v))

		if err != nil {
			l.Errorf("SaveOrderApp error: ", " | ", err)
			return nil, err
		}

		if v == nil {
			voucherNullErr := errors.New("is not possible to save order because this voucher does not exist")
			l.Infof("SaveOrderApp output: ", " | ", voucherNullErr)
			return nil, voucherNullErr
		}

		voucher = *v
	}

	for _, orderItems := range order.Items {

		product, err := app.productUseCase.GetProductByID(orderItems.ProductID)

		if err != nil {
			l.Errorf("SaveOrderApp error: ", " | ", err)
			return nil, err
		}

		if product == nil {
			productNullErr := errors.New("is not possible to save order because this product does not exists null")
			l.Infof("SaveOrderApp output: ", " | ", productNullErr)
			return nil, productNullErr
		}

		pi := entity.ProductItem{
			ID:            product.ID,
			Name:          product.Name,
			Quantity:      orderItems.Quantity,
			Category:      product.Category,
			Image:         product.Image,
			Description:   product.Description,
			Price:         product.Price,
			CreatedAt:     product.CreatedAt,
			DeactivatedAt: product.DeactivatedAt,
		}

		productList = append(productList, pi)

		tPrice := getTotalPrice(orderItems.Quantity, product.Price)

		totalPrice = totalPrice + tPrice

		var discount float64

		if voucher.Porcentage > 0 {
			discount = calculateDiscountByPercentage(voucher.Porcentage, tPrice)
		}

		opIn := dto.RequestOrderProduct{
			Quantity:   orderItems.Quantity,
			TotalPrice: tPrice,
			OrderID:    o.ID,
			ProductID:  &orderItems.ProductID,
			Discount:   discount,
		}

		opService, err := app.orderProductUseCase.SaveOrderProduct(opIn)

		if err != nil {
			l.Errorf("SaveOrderApp error: ", " | ", err)
			return nil, err
		}

		if opService == nil {
			orderProductNullErr := errors.New("is not possible to save order because it's null")
			l.Infof("SaveOrderApp output: ", " | ", orderProductNullErr)
			return nil, orderProductNullErr
		}
	}

	if voucher.Porcentage > 0 {
		totalPrice = calculateDiscountByPercentage(voucher.Porcentage, totalPrice)
	}

	outClient := dto.OutputClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}

	outOrder := &dto.OutputOrder{
		ID:               o.ID,
		Client:           outClient,
		TotalPrice:       totalPrice,
		Products:         productList,
		VoucherID:        o.VoucherID,
		Status:           o.Status,
		VerificationCode: o.VerificationCode,
		CreatedAt:        o.CreatedAt,
	}

	l.Infof("SaveOrderApp output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

// ========== Product ==========

func (app hermesFoodsApp) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("SaveProduct: ", " | ", ps.MarshalString(product))

	pRepo, err := app.productUseCase.SaveProduct(product)

	if err != nil {
		l.Errorf("SaveProduct error: ", " | ", err)
		return nil, err
	}

	if pRepo == nil {
		l.Infof("SaveProduct output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	out := &dto.OutputProduct{
		ID:            pRepo.ID,
		Name:          pRepo.Name,
		Category:      pRepo.Category,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt,
		DeactivatedAt: pRepo.DeactivatedAt,
	}

	l.Infof("SaveProduct output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	l.Infof("GetProductByCategory: ", " | ", category)
	productList := make([]dto.OutputProduct, 0)

	products, err := app.productUseCase.GetProductByCategory(category)

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
			Category:      products[i].Category,
			Image:         products[i].Image,
			Description:   products[i].Description,
			Price:         products[i].Price,
			CreatedAt:     products[i].CreatedAt,
			DeactivatedAt: products[i].CreatedAt,
		}
		productList = append(productList, product)
	}

	l.Infof("GetProductByCategory output: ", " | ", productList)
	return productList, nil
}

func (app hermesFoodsApp) UpdateProductByID(id int64, product dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("UpdateProductByID: ", " | ", id, " | ", ps.MarshalString(product))

	pByID, err := app.productUseCase.GetProductByID(id)

	if err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	if pByID == nil {
		productNullErr := errors.New("was not found any product with this id")
		l.Infof("UpdateProductByID output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	p, err := app.productUseCase.UpdateProductByID(id, product)

	if err != nil {
		l.Errorf("UpdateProductByID error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		productNullErr := errors.New("is not possible to save product because it's null")
		l.Errorf("UpdateProductByID output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	out := &dto.OutputProduct{
		ID:            p.ID,
		Name:          p.Name,
		Category:      p.Category,
		Image:         p.Image,
		Description:   p.Description,
		Price:         p.Price,
		CreatedAt:     p.CreatedAt,
		DeactivatedAt: p.DeactivatedAt,
	}

	l.Infof("UpdateProductByID output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app hermesFoodsApp) DeleteProductByID(id int64) error {
	l.Infof("DeleteProductByID: ", " | ", id)

	pByID, err := app.productUseCase.GetProductByID(id)

	if err != nil {
		l.Errorf("DeleteProductByID error: ", " | ", err)
		return err
	}

	if pByID == nil {
		productNullErr := errors.New("was not found any product with this id")
		l.Infof("DeleteProductByID output: ", " | ", productNullErr)
		return productNullErr
	}

	l.Infof("DeleteProductByID output: ", " | ", nil)
	return app.productUseCase.DeleteProductByID(id)
}

// ========== Voucher ==========

func (app hermesFoodsApp) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("SaveVoucher: ", " | ", ps.MarshalString(voucher))

	rVoucher, err := app.voucherUseCase.SaveVoucher(voucher)

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
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("SaveVoucher output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app hermesFoodsApp) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {
	l.Infof("GetVoucherByID: ", " | ", id)

	rVoucher, err := app.voucherUseCase.GetVoucherByID(id)

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
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("GetVoucherByID output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

func (app hermesFoodsApp) UpdateVoucherByID(id int64, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("UpdateVoucherByID: ", " | ", id, " | ", ps.MarshalString(voucher))

	rVoucher, err := app.voucherUseCase.UpdateVoucherByID(id, voucher)

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
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
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
