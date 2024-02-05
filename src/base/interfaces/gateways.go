package interfaces

import (
	"context"
	"fiap-hf-src/src/base/dto"
)

type PaymentGateway interface {
	DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error)
}

type ClientGateway interface {
	GetClientByID(id int64) (*dto.OutputClient, error)
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(client dto.RequestClient) (*dto.OutputClient, error)
}

type OrderGateway interface {
	SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error)
	UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error)
	GetOrders() ([]dto.OutputOrder, error)
	GetOrderByID(id int64) (*dto.OutputOrder, error)
}

type OrderProductGateway interface {
	GetAllOrderProduct() ([]dto.OutputOrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]dto.OutputOrderProduct, error)
	SaveOrderProduct(order dto.RequestOrderProduct) (*dto.OutputOrderProduct, error)
}

type ProductGateway interface {
	SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(id int64, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	GetProductByID(id int64) (*dto.OutputProduct, error)
	DeleteProductByID(id int64) error
}

type VoucherGateway interface {
	GetVoucherByID(id int64) (*dto.OutputVoucher, error)
	SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id int64, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}
