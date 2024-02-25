package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.VoucherUseCase = (*voucherUseCase)(nil)

type voucherUseCase struct {
	gateway interfaces.VoucherGateway
}

func NewVoucherUseCase(gateway interfaces.VoucherGateway) *voucherUseCase {
	return &voucherUseCase{gateway: gateway}
}

func (v *voucherUseCase) SaveVoucher(reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error) {

	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return nil, err
	}

	if len(voucher.Code) == 0 {
		return nil, errors.New("the voucher code is null or not valid")
	}

	if voucher.Porcentage < 0 || voucher.Porcentage > 101 {
		return nil, errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return v.gateway.SaveVoucher(reqVoucher)
}

func (v *voucherUseCase) UpdateVoucherByID(id int64, reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return nil, err
	}

	if len(voucher.Code) == 0 {
		return nil, errors.New("the voucher code is null or not valid")
	}

	if voucher.Porcentage < 0 || voucher.Porcentage > 101 {
		return nil, errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return v.gateway.UpdateVoucherByID(id, reqVoucher)
}

func (v *voucherUseCase) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}
	return v.gateway.GetVoucherByID(id)
}
