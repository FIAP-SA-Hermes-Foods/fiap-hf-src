package db

import (
	"context"
	"fiap-hf-src/src/base/interfaces"
	l "fiap-hf-src/src/base/logger"
	"fiap-hf-src/src/core/entity"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
)

var (
	queryGetVoucherByID    = `SELECT * FROM voucher WHERE id = $1`
	querySaveVoucher       = `INSERT INTO voucher (id, code, percentage, created_at, expires_at) VALUES (DEFAULT, $1, $2, now(), $3) RETURNING id, expires_at`
	queryUpdateVoucherByID = `UPDATE voucher SET code = $1, percentage = $2, expires_at = $3 WHERE id = $4 RETURNING id, expires_at`
)

var _ interfaces.VoucherDB = (*voucherDB)(nil)

type voucherDB struct {
	Ctx      context.Context
	Database interfaces.SQLDatabase
}

func NewVoucherDB(ctx context.Context, db interfaces.SQLDatabase) *voucherDB {
	return &voucherDB{Ctx: ctx, Database: db}
}

func (v *voucherDB) GetVoucherByID(id int64) (*entity.Voucher, error) {
	l.Infof("GetVoucherByID received input: ", " | ", id)
	if err := v.Database.Connect(); err != nil {
		l.Errorf("GetVoucherByID connect error: ", " | ", err)
		return nil, err
	}

	defer v.Database.Close()

	var outVoucher = new(entity.Voucher)

	if err := v.Database.Query(queryGetVoucherByID, id); err != nil {
		l.Errorf("GetVoucherByID query error: ", " | ", err)
		return nil, err
	}

	for v.Database.GetNextRows() {
		err := v.Database.Scan(
			&outVoucher.ID,
			&outVoucher.Code,
			&outVoucher.Porcentage,
			&outVoucher.CreatedAt.Value,
			&outVoucher.ExpiresAt.Value,
		)

		if err != nil {
			l.Errorf("GetVoucherByID scan error: ", " | ", err)
			return nil, err
		}
	}

	l.Infof("GetVoucherByID output: ", " | ", ps.MarshalString(outVoucher))
	return outVoucher, nil
}

func (v *voucherDB) SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error) {
	l.Infof("SaveVoucher received input: ", " | ", voucher)
	if err := v.Database.Connect(); err != nil {
		l.Errorf("SaveVoucher connect error: ", " | ", err)
		return nil, err
	}

	defer v.Database.Close()

	if err := v.Database.PrepareStmt(querySaveVoucher); err != nil {
		l.Errorf("SaveVoucher prepare error: ", " | ", err)
		return nil, err
	}

	defer v.Database.CloseStmt()

	var outVoucher = &entity.Voucher{
		Code:       voucher.Code,
		Porcentage: voucher.Porcentage,
		ExpiresAt: com.ExpiresAt{
			Value: voucher.ExpiresAt.Value,
		},
	}

	v.Database.QueryRow(voucher.Code, voucher.Porcentage, voucher.ExpiresAt)

	if err := v.Database.ScanStmt(&outVoucher.ID, &outVoucher.CreatedAt.Value); err != nil {
		l.Errorf("SaveVoucher scan error: ", " | ", err)
		return nil, err
	}
	l.Infof("SaveVoucher output: ", " | ", ps.MarshalString(outVoucher))
	return outVoucher, nil
}

func (v *voucherDB) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
	l.Infof("UpdateVoucherByID received input: ", " | ", id, " | ", voucher)
	if err := v.Database.Connect(); err != nil {
		l.Errorf("UpdateVoucherByID connect error: ", " | ", err)
		return nil, err
	}

	defer v.Database.Close()

	if err := v.Database.PrepareStmt(queryUpdateVoucherByID); err != nil {
		l.Errorf("UpdateVoucherByID prepare error: ", " | ", err)
		return nil, err
	}

	defer v.Database.CloseStmt()

	var outVoucher = &entity.Voucher{
		Code:       voucher.Code,
		Porcentage: voucher.Porcentage,
		CreatedAt: com.CreatedAt{
			Value: voucher.CreatedAt.Value,
		},
		ExpiresAt: com.ExpiresAt{
			Value: &voucher.CreatedAt.Value,
		},
	}

	v.Database.QueryRow(voucher.Code, voucher.Porcentage, voucher.ExpiresAt.Value, id)

	if err := v.Database.ScanStmt(&outVoucher.ID, &outVoucher.CreatedAt.Value); err != nil {
		l.Errorf("UpdateVoucherByID scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("UpdateVoucherByID output: ", " | ", ps.MarshalString(outVoucher))
	return outVoucher, nil
}
