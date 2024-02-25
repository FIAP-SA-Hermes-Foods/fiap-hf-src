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

// go test -v -count=1 -failfast -cover -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		product entity.Product
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Product
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			repo := NewProductDB(tc.ctx, tc.mockDB)

			out, err := repo.SaveProduct(tc.args.product)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetProductByID$
func Test_GetProductByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Product
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			repo := NewProductDB(tc.ctx, tc.mockDB)

			out, err := repo.GetProductByID(tc.args.id)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		category string
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Product
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			repo := NewProductDB(tc.ctx, tc.mockDB)

			_, err := repo.GetProductByCategory(tc.args.category)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_UpdateProductByID$
func Test_UpdateProductByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id      int64
		product entity.Product
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Product
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			repo := NewProductDB(tc.ctx, tc.mockDB)

			out, err := repo.UpdateProductByID(tc.args.id, tc.args.product)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_DeleteProductByID$
func Test_DeleteProductByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id      int64
		product entity.Product
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Product
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			WantOutput: &entity.Product{},
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
			repo := NewProductDB(tc.ctx, tc.mockDB)

			err := repo.DeleteProductByID(tc.args.id)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}
		})
	}
}
