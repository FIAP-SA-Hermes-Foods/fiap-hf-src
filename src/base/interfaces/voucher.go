package interfaces

import "fiap-hf-src/src/core/entity"

type VoucherRepository interface {
	GetVoucherByID(id int64) (*entity.Voucher, error)
	SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error)
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error)
}

type VoucherService interface {
	SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error)
	GetVoucherByID(id int64) error
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error)
}
