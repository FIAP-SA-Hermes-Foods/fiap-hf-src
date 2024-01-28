package interfaces

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/core/entity"
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
