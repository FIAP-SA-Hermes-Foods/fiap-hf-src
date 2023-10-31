package voucher

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -cover -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {
	ctx := context.Background()

	type args struct {
		product entity.Voucher
	}

	valExpiresAt := valueObject.ExpiresAt{}
	valExpiresAt.SetTimeFromString("0001-01-01T00:00:00Z")

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Voucher
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Voucher{
				ID:         0,
				Code:       "",
				Porcentage: 0,
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: valExpiresAt,
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name:       "connection_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name:       "prepare_stmt_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errPrepareStmt"),
			},

			isWantError: true,
		},
		{
			name:       "prepare_stmt_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name:       "error_scan_stmt",
			args:       args{},
			ctx:        nil,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult:   nil,
				WantRows:     &sql.Rows{},
				WantErr:      errors.New("errScanStmt"),
				WantNextRows: false,
			},
			isWantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewVoucherRepository(tc.ctx, tc.mockDB)

			out, err := repo.SaveVoucher(tc.args.product)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetVoucherByID$
func Test_GetVoucherByID(t *testing.T) {
	ctx := context.Background()

	type args struct {
		id int64
	}

	valExpiresAt := valueObject.ExpiresAt{}
	valExpiresAt.SetTimeFromString("0001-01-01T00:00:00Z")

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Voucher
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Voucher{
				ID:         0,
				Code:       "",
				Porcentage: 0,
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: valExpiresAt,
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name:       "connection_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name:       "query_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errQuery"),
			},

			isWantError: true,
		},
		{
			name:       "prepare_stmt_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errPrepareStmt"),
			},

			isWantError: true,
		},
		{
			name:       "prepare_stmt_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name:       "error_scan_stmt",
			args:       args{},
			ctx:        nil,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult:   nil,
				WantRows:     &sql.Rows{},
				WantErr:      errors.New("errScanStmt"),
				WantNextRows: false,
			},
			isWantError: true,
		},
		{
			name: "error_scan",
			args: args{},
			ctx:  nil,
			mockDB: &mockDb{
				WantResult:   nil,
				WantRows:     &sql.Rows{},
				WantErr:      errors.New("errScan"),
				WantNextRows: true,
			},
			isWantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewVoucherRepository(tc.ctx, tc.mockDB)

			out, err := repo.GetVoucherByID(tc.args.id)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_UpdateVoucherByID$
func Test_UpdateVoucherByID(t *testing.T) {
	ctx := context.Background()

	type args struct {
		id      int64
		product entity.Voucher
	}

	tt, _ := time.Parse("02-01-2006 15:04:05", "0001-01-01T00:00:00Z")

	valExpiresAt := valueObject.ExpiresAt{Value: &tt}
	// 02-01-2006 15:04:05
	// err := valExpiresAt.SetTimeFromString("01-01-0001 00:00:00")
	// log.Print(err)

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Voucher
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Voucher{
				ID:         0,
				Code:       "",
				Porcentage: 0,
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: valExpiresAt,
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name:       "connection_error",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name: "prepare_stmt_error",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Voucher{
				ID:         0,
				Code:       "",
				Porcentage: 0,
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: valExpiresAt,
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errPrepareStmt"),
			},

			isWantError: true,
		},
		{
			name: "prepare_stmt_error",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Voucher{
				ID:         0,
				Code:       "",
				Porcentage: 0,
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: valExpiresAt,
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name:       "error_scan_stmt",
			args:       args{},
			ctx:        nil,
			WantOutput: &entity.Voucher{},
			mockDB: &mockDb{
				WantResult:   nil,
				WantRows:     &sql.Rows{},
				WantErr:      errors.New("errScanStmt"),
				WantNextRows: false,
			},
			isWantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewVoucherRepository(tc.ctx, tc.mockDB)

			out, err := repo.UpdateVoucherByID(tc.args.id, tc.args.product)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}
		})
	}
}
