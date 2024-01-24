package voucher

import (
	"context"
	"fiap-hf-src/internal/core/entity"
	com "fiap-hf-src/internal/core/entity/common"
	"fiap-hf-src/internal/core/useCase/db"
	l "fiap-hf-src/pkg/logger"
)

var (
	queryGetVoucherByID    = `SELECT * FROM voucher WHERE id = $1`
	querySaveVoucher       = `INSERT INTO voucher (id, code, percentage, created_at, expires_at) VALUES (DEFAULT, $1, $2, now(), $3) RETURNING id, expires_at`
	queryUpdateVoucherByID = `UPDATE voucher SET code = $1, percentage = $2, expires_at = $3 WHERE id = $4 RETURNING id, expires_at`
)

type VoucherRepository interface {
	GetVoucherByID(id int64) (*entity.Voucher, error)
	SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error)
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error)
}

type voucherRepository struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewVoucherRepository(ctx context.Context, db db.SQLDatabase) VoucherRepository {
	return voucherRepository{Ctx: ctx, Database: db}
}

func (v voucherRepository) GetVoucherByID(id int64) (*entity.Voucher, error) {
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

	l.Infof("GetVoucherByID output: ", " | ", outVoucher.MarshalString())
	return outVoucher, nil
}

func (v voucherRepository) SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error) {
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
	l.Infof("SaveVoucher output: ", " | ", outVoucher.MarshalString())
	return outVoucher, nil
}

func (v voucherRepository) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
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

	l.Infof("UpdateVoucherByID output: ", " | ", outVoucher.MarshalString())
	return outVoucher, nil
}
