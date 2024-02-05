package mocks

import (
	"fiap-hf-src/src/base/dto"
	"strings"
)

type ClientGatewayMock struct {
	WantOut *dto.OutputClient
	WantErr error
}

func (g ClientGatewayMock) GetClientByID(id int64) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetClientByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ClientGatewayMock) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetClientByCPF") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ClientGatewayMock) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveClient") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

type OrderGatewayMock struct {
	WantOut     *dto.OutputOrder
	WantOutList []dto.OutputOrder
	WantErr     error
}

func (g OrderGatewayMock) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveOrder") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g OrderGatewayMock) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateOrderByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g OrderGatewayMock) GetOrders() ([]dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetOrders") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderGatewayMock) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetOrderByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

type OrderProductGatewayMock struct {
	WantOut     *dto.OutputOrderProduct
	WantOutList []dto.OutputOrderProduct
	WantErr     error
}

func (g OrderProductGatewayMock) GetAllOrderProduct() ([]dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetAllOrderProduct") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderProductGatewayMock) GetAllOrderProductByOrderID(id int64) ([]dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetAllOrderProductByOrderID") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderProductGatewayMock) SaveOrderProduct(order dto.RequestOrderProduct) (*dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveOrderProduct") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

// ====== Prtoduct =======
type ProductGatewayMock struct {
	WantOut     *dto.OutputProduct
	WantOutList []dto.OutputProduct
	WantErr     error
}

func (g ProductGatewayMock) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveProduct") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductGatewayMock) UpdateProductByID(id int64, product dto.RequestProduct) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateProductByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductGatewayMock) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetProductByCategory") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g ProductGatewayMock) GetProductByID(id int64) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetProductByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductGatewayMock) DeleteProductByID(id int64) error {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errDeleteProductByID") {
		return g.WantErr
	}
	return nil
}

// ====== Voucher =======
type VoucherGateway struct {
	WantOut *dto.OutputVoucher
	WantErr error
}

func (g VoucherGateway) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetVoucherByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g VoucherGateway) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveVoucher") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g VoucherGateway) UpdateVoucherByID(id int64, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateVoucherByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}
