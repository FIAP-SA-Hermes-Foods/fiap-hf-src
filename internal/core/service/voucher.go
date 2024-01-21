package service

import (
	"errors"
	"fiap-hf-src/internal/core/entity"
)

type VoucherService interface {
	SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error)
	GetVoucherByID(id int64) error
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error)
}

type voucherService struct {
	Voucher *entity.Voucher
}

func NewVoucherService(voucher *entity.Voucher) VoucherService {
	if voucher == nil {
		return voucherService{Voucher: new(entity.Voucher)}
	}
	return voucherService{Voucher: voucher}
}

func (o voucherService) SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error) {

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return nil, err
	}

	if len(voucher.Code) == 0 {
		return nil, errors.New("the voucher code is null or not valid")
	}

	if voucher.Porcentage < 0 || voucher.Porcentage > 101 {
		return nil, errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return &voucher, nil
}

func (o voucherService) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return nil, err
	}

	if len(voucher.Code) == 0 {
		return nil, errors.New("the voucher code is null or not valid")
	}

	if voucher.Porcentage < 0 || voucher.Porcentage > 101 {
		return nil, errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return &voucher, nil
}

func (c voucherService) GetVoucherByID(id int64) error {
	if id < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}
