package db

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/src/base/mocks"
	"fiap-hf-src/src/core/entity"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -count=1 -failfast -cover -run ^Test_GetOrders$
func Test_GetOrders(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Order
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			repo := NewOrderDB(tc.ctx, tc.mockDB)

			_, err := repo.GetOrders()

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetOrderByID$
func Test_GetOrderByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Order
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			repo := NewOrderDB(tc.ctx, tc.mockDB)

			out, err := repo.GetOrderByID(tc.args.id)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		order entity.Order
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Order
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			repo := NewOrderDB(tc.ctx, tc.mockDB)

			out, err := repo.SaveOrder(tc.args.order)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {

				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_UpdateOrderByID$
func Test_UpdateOrderByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id    int64
		order entity.Order
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Order
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Order{},
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			mockDB: &mocks.MockDb{
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
			repo := NewOrderDB(tc.ctx, tc.mockDB)

			out, err := repo.UpdateOrderByID(tc.args.id, tc.args.order)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}
