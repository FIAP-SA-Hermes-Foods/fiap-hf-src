package voucher

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
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
	Database psqldb.PostgresDB
}

func NewVoucherRepository(ctx context.Context, db psqldb.PostgresDB) VoucherRepository {
	return voucherRepository{Ctx: ctx, Database: db}
}

func (v voucherRepository) GetVoucherByID(id int64) (*entity.Voucher, error) {
	if err := v.Database.Connect(); err != nil {
		return nil, err
	}

	defer v.Database.Close()

	var outVoucher = new(entity.Voucher)

	if err := v.Database.Query(queryGetVoucherByID, id); err != nil {
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
			return nil, err
		}
	}

	return outVoucher, nil
}

func (v voucherRepository) SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error) {
	if err := v.Database.Connect(); err != nil {
		return nil, err
	}

	defer v.Database.Close()

	if err := v.Database.PrepareStmt(querySaveVoucher); err != nil {
		return nil, err
	}

	defer v.Database.CloseStmt()

	var outVoucher = &entity.Voucher{
		Code:       voucher.Code,
		Porcentage: voucher.Porcentage,
		ExpiresAt: valueObject.ExpiresAt{
			Value: voucher.ExpiresAt.Value,
		},
	}

	v.Database.QueryRow(voucher.Code, voucher.Porcentage, voucher.ExpiresAt)

	if err := v.Database.ScanStmt(&outVoucher.ID, &outVoucher.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outVoucher, nil
}

func (v voucherRepository) UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
	if err := v.Database.Connect(); err != nil {
		return nil, err
	}

	defer v.Database.Close()

	if err := v.Database.PrepareStmt(queryUpdateVoucherByID); err != nil {
		return nil, err
	}

	defer v.Database.CloseStmt()

	var outVoucher = &entity.Voucher{
		Code:       voucher.Code,
		Porcentage: voucher.Porcentage,
		CreatedAt: valueObject.CreatedAt{
			Value: voucher.CreatedAt.Value,
		},
		ExpiresAt: valueObject.ExpiresAt{
			Value: &voucher.CreatedAt.Value,
		},
	}

	v.Database.QueryRow(voucher.Code, voucher.Porcentage, voucher.ExpiresAt.Value, id)

	if err := v.Database.ScanStmt(&outVoucher.ID, &outVoucher.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outVoucher, nil
}
