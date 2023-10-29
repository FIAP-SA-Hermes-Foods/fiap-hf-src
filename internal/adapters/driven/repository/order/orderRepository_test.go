package order

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	"testing"
)

// go test -v -count=1 -failfast -cover -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {
	ctx := context.Background()

	type args struct {
		order entity.Order
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Order
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Order{},
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
			WantOutput: &entity.Order{},
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
			WantOutput: &entity.Order{},
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
			WantOutput: &entity.Order{},
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
			WantOutput: &entity.Order{},
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
			repo := NewOrderRepository(tc.ctx, tc.mockDB)

			out, err := repo.SaveOrder(tc.args.order)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}

		})
	}
}
