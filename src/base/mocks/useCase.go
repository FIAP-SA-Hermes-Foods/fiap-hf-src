package mocks

import (
	"fiap-hf-src/src/base/dto"
	"strings"
)

type ClientUseCaseMock struct {
	WantOut *dto.OutputClient
	WantErr error
}

func (g ClientUseCaseMock) GetClientByID(id int64) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetClientByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ClientUseCaseMock) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetClientByCPF") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ClientUseCaseMock) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveClient") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

type OrderUseCaseMock struct {
	WantOut     *dto.OutputOrder
	WantOutList []dto.OutputOrder
	WantErr     error
}

func (g OrderUseCaseMock) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveOrder") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g OrderUseCaseMock) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateOrderByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g OrderUseCaseMock) GetOrders() ([]dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetOrders") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderUseCaseMock) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetOrderByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

type OrderProductUseCaseMock struct {
	WantOut     *dto.OutputOrderProduct
	WantOutList []dto.OutputOrderProduct
	WantErr     error
}

func (g OrderProductUseCaseMock) GetAllOrderProduct() ([]dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetAllOrderProduct") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderProductUseCaseMock) GetAllOrderProductByOrderID(id int64) ([]dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetAllOrderProductByOrderID") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g OrderProductUseCaseMock) SaveOrderProduct(order dto.RequestOrderProduct) (*dto.OutputOrderProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveOrderProduct") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

// ====== Prtoduct =======
type ProductUseCaseMock struct {
	WantOut     *dto.OutputProduct
	WantOutList []dto.OutputProduct
	WantErr     error
}

func (g ProductUseCaseMock) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveProduct") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductUseCaseMock) UpdateProductByID(id int64, product dto.RequestProduct) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateProductByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductUseCaseMock) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetProductByCategory") {
		return nil, g.WantErr
	}
	return g.WantOutList, nil
}

func (g ProductUseCaseMock) GetProductByID(id int64) (*dto.OutputProduct, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetProductByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g ProductUseCaseMock) DeleteProductByID(id int64) error {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errDeleteProductByID") {
		return g.WantErr
	}
	return nil
}

// ====== Voucher =======
type VoucherUseCase struct {
	WantOut *dto.OutputVoucher
	WantErr error
}

func (g VoucherUseCase) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errGetVoucherByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g VoucherUseCase) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errSaveVoucher") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}

func (g VoucherUseCase) UpdateVoucherByID(id int64, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if g.WantErr != nil && strings.EqualFold(g.WantErr.Error(), "errUpdateVoucherByID") {
		return nil, g.WantErr
	}
	return g.WantOut, nil
}
