package db

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.VoucherGateway = (*voucherGateway)(nil)

type voucherGateway struct {
	db interfaces.VoucherDB
}

func NewVoucherGateway(db interfaces.VoucherDB) *voucherGateway {
	return &voucherGateway{db: db}
}

func (v *voucherGateway) GetVoucherByID(id int64) (*dto.OutputVoucher, error) {

	outDB, err := v.db.GetVoucherByID(id)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputVoucher{
		ID:         outDB.ID,
		Code:       outDB.Code,
		Porcentage: outDB.Porcentage,
		CreatedAt:  outDB.CreatedAt.Format(),
		ExpiresAt:  outDB.ExpiresAt.Format(),
	}
	return out, nil
}

func (v *voucherGateway) SaveVoucher(reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error) {

	voucher := reqVoucher.Voucher()

	outDB, err := v.db.SaveVoucher(voucher)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputVoucher{
		ID:         outDB.ID,
		Code:       outDB.Code,
		Porcentage: outDB.Porcentage,
		CreatedAt:  outDB.CreatedAt.Format(),
		ExpiresAt:  outDB.ExpiresAt.Format(),
	}

	return out, nil
}

func (v *voucherGateway) UpdateVoucherByID(id int64, reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	voucher := reqVoucher.Voucher()

	outDB, err := v.db.SaveVoucher(voucher)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputVoucher{
		ID:         outDB.ID,
		Code:       outDB.Code,
		Porcentage: outDB.Porcentage,
		CreatedAt:  outDB.CreatedAt.Format(),
		ExpiresAt:  outDB.ExpiresAt.Format(),
	}

	return out, nil
}
