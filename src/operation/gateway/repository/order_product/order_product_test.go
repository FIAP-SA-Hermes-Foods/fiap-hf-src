package orderproduct

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/src/core/entity"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"fmt"
	"log"
	"os/exec"
	"testing"
)

// go test -v -count=1 -failfast -cover -run ^Test_GetAllOrderProduct$
func Test_GetAllOrderProduct(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.OrderProduct
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			repo := NewOrderProductRepository(tc.ctx, tc.mockDB)

			_, err := repo.GetAllOrderProduct()

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetAllOrderProductByOrderID$
func Test_GetAllOrderProductByOrderID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.OrderProduct
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			repo := NewOrderProductRepository(tc.ctx, tc.mockDB)

			_, err := repo.GetAllOrderProductByOrderID(tc.args.id)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_SaveOrderProduct$
func Test_SaveOrderProduct(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		order entity.OrderProduct
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.OrderProduct
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name:       "success",
			args:       args{},
			ctx:        ctx,
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			WantOutput: &entity.OrderProduct{},
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
			repo := NewOrderProductRepository(tc.ctx, tc.mockDB)

			out, err := repo.SaveOrderProduct(tc.args.order)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}

func printArt() {
	for i := range goTestArt {
		b, err := exec.Command("echo", "-ne", goTestArt[i]).Output()
		if err != nil {
			log.Printf("error to print the art -> %v", err)
		}

		fmt.Printf("%s", b)

	}
}

var goTestArt = []string{
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;244;245;245m@\x1b[0m\x1b[38;2;64;75;90m~\x1b[0m\x1b[38;2;51;62;78m-\x1b[0m\x1b[38;2;48;60;75m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;206;209;213mW\x1b[0m\x1b[38;2;137;196;234m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;60;71;86m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;128;128;128m7\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;185;189;194m$\x1b[0m\x1b[38;2;91;102;114m=\x1b[0m\x1b[38;2;60;73;88m~\x1b[0m\x1b[38;2;53;65;81m~\x1b[0m\x1b[38;2;57;69;85m~\x1b[0m\x1b[38;2;83;94;107m+\x1b[0m\x1b[38;2;166;172;178m9\x1b[0m\x1b[38;2;253;253;253mÑ\x1b[0m\x1b[38;2;106;146;175m8\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;57;71;86m~\x1b[0m\x1b[38;2;104;148;176m8\x1b[0m\x1b[38;2;68;89;106m+\x1b[0m\x1b[38;2;153;218;255mW\x1b[0m\x1b[38;2;251;252;252mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;122;131;141m7\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;139;199;237m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;151;214;255mW\x1b[0m\x1b[38;2;50;61;74m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;127;136;145m7\x1b[0m\x1b[38;2;63;79;96m~\x1b[0m\x1b[38;2;139;199;237m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;195;232m$\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;152;217;255mW\x1b[0m\x1b[38;2;63;79;95m~\x1b[0m\x1b[38;2;89;120;143m=\x1b[0m\x1b[38;2;151;216;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;53;68;82m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;253;253;253mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;83;110;133m=\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;144;206;245mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;51;60;75m-\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;53;66;80m~\x1b[0m\x1b[38;2;98;110;121m=\x1b[0m\x1b[38;2;182;187;191m$\x1b[0m\x1b[38;2;100;140;164m7\x1b[0m\x1b[38;2;130;182;208m9\x1b[0m\x1b[38;2;196;200;204mW\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;243;244;245m@\x1b[0m\x1b[38;2;140;197;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;78;102;123m=\x1b[0m\x1b[38;2;95;135;162m7\x1b[0m\x1b[38;2;95;135;162m7\x1b[0m\x1b[38;2;64;82;101m~\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;141;201;240mW\x1b[0m\x1b[38;2;151;216;255mW\x1b[0m\x1b[38;2;154;220;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;129;181;207m9\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;192;195;201mW\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;148;212;246mW\x1b[0m\x1b[38;2;155;221;255mW\x1b[0m\x1b[38;2;181;185;191m$\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;26;26;26m.\x1b[0m\x1b[38;2;164;164;164m9\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;236;237;239m@\x1b[0m\x1b[38;2;139;199;236m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;140;198;237m$\x1b[0m\x1b[38;2;85;113;136m=\x1b[0m\x1b[38;2;137;194;232m$\x1b[0m\x1b[38;2;141;200;238mW\x1b[0m\x1b[38;2;51;60;72m-\x1b[0m\x1b[38;2;140;200;239m$\x1b[0m\x1b[38;2;145;207;246mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;154;219;255mW\x1b[0m\x1b[38;2;104;144;167m8\x1b[0m\x1b[38;2;54;66;80m~\x1b[0m\x1b[38;2;53;63;77m-\x1b[0m\x1b[38;2;54;64;79m~\x1b[0m\x1b[38;2;51;61;75m-\x1b[0m\x1b[38;2;92;124;146m7\x1b[0m\x1b[38;2;155;222;255m#\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;50;62;78m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;251;251;252mÑ\x1b[0m\x1b[38;2;179;184;189m$\x1b[0m\x1b[38;2;195;199;203mW\x1b[0m\x1b[38;2;253;253;253mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;254;254;254mÑ\x1b[0m\x1b[38;2;136;194;230m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;138;197;234m$\x1b[0m\x1b[38;2;122;171;203m9\x1b[0m\x1b[38;2;132;188;224m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;148;212;251mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;69;89;107m+\x1b[0m\x1b[38;2;60;73;88m~\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;245;245;245m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;252;252;252mÑ\x1b[0m\x1b[38;2;52;62;76m-\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;158;164;171m9\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;54;67;82m~\x1b[0m\x1b[38;2;252;252;252mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;250;250;250mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;66;84;101m+\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;136;194;230m$\x1b[0m\x1b[38;2;56;70;84m~\x1b[0m\x1b[38;2;239;238;238m@\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;251;251;251mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;54;67;82m~\x1b[0m\x1b[38;2;152;217;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;61;78;93m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;61;73;89m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;103;111;124m=\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;83;94;108m+\x1b[0m\x1b[38;2;225;227;230m#\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;228;230;232m@\x1b[0m\x1b[38;2;139;199;236m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;84;112;134m=\x1b[0m\x1b[38;2;103;111;119m=\x1b[0m\x1b[38;2;54;71;88m~\x1b[0m\x1b[38;2;150;215;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;94;128;152m7\x1b[0m\x1b[38;2;227;228;229m#\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;247;247;247mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;83;111;129m=\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;131;185;213m$\x1b[0m\x1b[38;2;244;245;246m@\x1b[0m\x1b[38;2;250;251;251mÑ\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;126;134;143m7\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;217;219;222m#\x1b[0m\x1b[38;2;137;197;234m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;32;40;48m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;231;232;234m@\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;137;195;232m$\x1b[0m\x1b[38;2;142;148;156m8\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;253;253;253mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;87;97;112m+\x1b[0m\x1b[38;2;134;141;151m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;249;249;249mÑ\x1b[0m\x1b[38;2;145;205;237mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;154;221;255mW\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;77;89;103m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;252;252;252mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;115;161;192m9\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;82;109;131m=\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;32;40;46m,\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;104;146;171m8\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;90;121;145m=\x1b[0m\x1b[38;2;239;238;238m@\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;216;218;221m#\x1b[0m\x1b[38;2;233;234;237m@\x1b[0m\x1b[38;2;58;68;86m~\x1b[0m\x1b[38;2;87;97;111m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;59;72;88m~\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;86;114;134m=\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;65;77;91m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;236;236;236m@\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;23;23;23m.\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;214;217;220m#\x1b[0m\x1b[38;2;139;199;236m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;136;195;232m$\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;154;219;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;136;195;231m$\x1b[0m\x1b[38;2;116;125;135m7\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;248;248;248mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;192;195;201mW\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;68;85m~\x1b[0m\x1b[38;2;57;69;85m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;96;132;152m7\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;143;151;159m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;234;233;233m@\x1b[0m\x1b[38;2;248;248;249mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;28;28;28m.\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;232;233;235m@\x1b[0m\x1b[38;2;141;202;239mW\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;67;77;90m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;30;37;44m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;45m,\x1b[0m\x1b[38;2;149;213;252mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;100;138;165m7\x1b[0m\x1b[38;2;177;181;186m$\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;241;241;241m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;55;68;83m~\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;130;182;210m9\x1b[0m\x1b[38;2;56;68;83m~\x1b[0m\x1b[38;2;57;70;85m~\x1b[0m\x1b[38;2;57;70;85m~\x1b[0m\x1b[38;2;231;231;230m@\x1b[0m\x1b[38;2;232;232;232m@\x1b[0m\x1b[38;2;49;63;79m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;18;18;18m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;18;18;18m.\x1b[0m\x1b[38;2;246;246;246mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;166;170;177m9\x1b[0m\x1b[38;2;142;202;241mW\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;195;233m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;139;199;236m$\x1b[0m\x1b[38;2;55;67;81m~\x1b[0m\x1b[38;2;187;190;194m$\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;243;243;243m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;138;145;155m8\x1b[0m\x1b[38;2;68;88;103m+\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;59;73;88m~\x1b[0m\x1b[38;2;226;208;204m#\x1b[0m\x1b[38;2;225;208;203m#\x1b[0m\x1b[38;2;225;208;203m#\x1b[0m\x1b[38;2;225;208;203m#\x1b[0m\x1b[38;2;225;208;203m#\x1b[0m\x1b[38;2;126;126;133m7\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;245;245;245m@\x1b[0m\x1b[38;2;31;31;31m,\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;36;36;36m,\x1b[0m\x1b[38;2;247;247;247mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;177;183;188m$\x1b[0m\x1b[38;2;133;190;226m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;138;198;234m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;138;199;236m$\x1b[0m\x1b[38;2;57;70;85m~\x1b[0m\x1b[38;2;210;212;215m#\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;233;233;233m@\x1b[0m\x1b[38;2;249;249;249mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;246;245;244m@\x1b[0m\x1b[38;2;55;69;84m~\x1b[0m\x1b[38;2;152;217;249mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;141;201;238mW\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;103;112;124m=\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;117;117;117m=\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;76;69;64m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;61;70;81m~\x1b[0m\x1b[38;2;56;69;83m~\x1b[0m\x1b[38;2;143;204;242mW\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;195;233m$\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;141;200;238mW\x1b[0m\x1b[38;2;135;193;230m$\x1b[0m\x1b[38;2;137;195;232m$\x1b[0m\x1b[38;2;141;201;240mW\x1b[0m\x1b[38;2;136;195;232m$\x1b[0m\x1b[38;2;138;197;235m$\x1b[0m\x1b[38;2;150;216;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;74;85;99m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;246;246;246mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;51;64;80m~\x1b[0m\x1b[38;2;228;217;187m#\x1b[0m\x1b[38;2;131;138;148m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;201;205;209mW\x1b[0m\x1b[38;2;64;85;103m+\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;76;71;65m~\x1b[0m\x1b[38;2;56;70;86m~\x1b[0m\x1b[38;2;67;85;103m+\x1b[0m\x1b[38;2;139;199;237m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;195;232m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;234;236;237m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;45;60;76m-\x1b[0m\x1b[38;2;227;216;186mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;84;93;101m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;198;198;198mW\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;138;145;154m8\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;220;209;182mW\x1b[0m\x1b[38;2;142;203;241mW\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;89;121;145m=\x1b[0m\x1b[38;2;78;71;64m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;56;70;85m~\x1b[0m\x1b[38;2;139;198;236m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;147;210;247mW\x1b[0m\x1b[38;2;74;95;114m+\x1b[0m\x1b[38;2;60;75;91m~\x1b[0m\x1b[38;2;141;199;229m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;97;131;154m7\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;233;221;190m#\x1b[0m\x1b[38;2;225;213;184mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;73;84;94m+\x1b[0m\x1b[38;2;222;211;183mW\x1b[0m\x1b[38;2;209;199;172mW\x1b[0m\x1b[38;2;209;199;172mW\x1b[0m\x1b[38;2;53;66;81m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;104;114;124m=\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;93;127;151m7\x1b[0m\x1b[38;2;68;86;103m+\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;65;71;78m~\x1b[0m\x1b[38;2;168;163;147m9\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;80;92;106m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;238;239;240m@\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;114;119;118m=\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;101;108;110m=\x1b[0m\x1b[38;2;211;201;173mW\x1b[0m\x1b[38;2;62;74;88m~\x1b[0m\x1b[38;2;213;203;174mW\x1b[0m\x1b[38;2;214;204;175mW\x1b[0m\x1b[38;2;215;205;175mW\x1b[0m\x1b[38;2;217;206;176mW\x1b[0m\x1b[38;2;51;63;79m~\x1b[0m\x1b[38;2;253;253;253mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;9;9;9m \x1b[0m\x1b[38;2;25;25;25m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;244;245;246m@\x1b[0m\x1b[38;2;142;202;241mW\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;140;200;239m$\x1b[0m\x1b[38;2;101;140;167m7\x1b[0m\x1b[38;2;76;71;67m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;60;71;83m~\x1b[0m\x1b[38;2;201;192;166m$\x1b[0m\x1b[38;2;207;197;170m$\x1b[0m\x1b[38;2;94;101;106m+\x1b[0m\x1b[38;2;209;200;175mW\x1b[0m\x1b[38;2;88;97;104m+\x1b[0m\x1b[38;2;232;219;189m#\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;149;149;140m8\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;147;211;250mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;135;192;219m$\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;248;248;248mÑ\x1b[0m\x1b[38;2;242;242;242m@\x1b[0m\x1b[38;2;236;236;236m@\x1b[0m\x1b[38;2;234;234;234m@\x1b[0m\x1b[38;2;147;147;147m8\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;27;27;27m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;213;216;218m#\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;140;199;237m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;77;103;120m+\x1b[0m\x1b[38;2;83;111;134m=\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;78;71;65m~\x1b[0m\x1b[38;2;182;177;155m9\x1b[0m\x1b[38;2;208;198;170m$\x1b[0m\x1b[38;2;113;123;135m7\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;52;65;81m~\x1b[0m\x1b[38;2;224;214;185mW\x1b[0m\x1b[38;2;180;176;158m9\x1b[0m\x1b[38;2;55;69;85m~\x1b[0m\x1b[38;2;188;180;159m$\x1b[0m\x1b[38;2;209;198;171mW\x1b[0m\x1b[38;2;52;65;82m~\x1b[0m\x1b[38;2;129;183;217m$\x1b[0m\x1b[38;2;140;201;239mW\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;137;196;233m$\x1b[0m\x1b[38;2;136;195;232m$\x1b[0m\x1b[38;2;137;195;232m$\x1b[0m\x1b[38;2;138;197;235m$\x1b[0m\x1b[38;2;139;199;237m$\x1b[0m\x1b[38;2;140;201;239mW\x1b[0m\x1b[38;2;129;137;147m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;21;21;21m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;248;248;248mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;108;150;179m8\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;149;212;252mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;154;222;254mW\x1b[0m\x1b[38;2;110;151;178m8\x1b[0m\x1b[38;2;149;214;246mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;67;85;103m+\x1b[0m\x1b[38;2;77;102;123m+\x1b[0m\x1b[38;2;77;71;64m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;136;137;130m7\x1b[0m\x1b[38;2;210;200;172mW\x1b[0m\x1b[38;2;226;215;185mW\x1b[0m\x1b[38;2;232;219;189m#\x1b[0m\x1b[38;2;52;66;82m~\x1b[0m\x1b[38;2;78;71;64m~\x1b[0m\x1b[38;2;71;70;69m~\x1b[0m\x1b[38;2;101;138;165m7\x1b[0m\x1b[38;2;115;160;191m8\x1b[0m\x1b[38;2;78;71;64m~\x1b[0m\x1b[38;2;65;84;103m+\x1b[0m\x1b[38;2;67;85;104m+\x1b[0m\x1b[38;2;85;116;139m=\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;76;70;64m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;24;24;24m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;232;232;232m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;102;141;168m8\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;151;215;255mW\x1b[0m\x1b[38;2;136;194;223m$\x1b[0m\x1b[38;2;247;247;248mÑ\x1b[0m\x1b[38;2;31;39;47m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;34;43;53m,\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;189;193;197mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;148;210;241mW\x1b[0m\x1b[38;2;201;192;166m$\x1b[0m\x1b[38;2;34;42;50m,\x1b[0m\x1b[38;2;73;82;91m~\x1b[0m\x1b[38;2;230;218;188m#\x1b[0m\x1b[38;2;152;218;251mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;55;68;82m~\x1b[0m\x1b[38;2;117;163;195m9\x1b[0m\x1b[38;2;64;67;72m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;131;128;123m7\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;81;81;81m~\x1b[0m\x1b[38;2;32;32;32m,\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;176;180;186m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;150;215;255mW\x1b[0m\x1b[38;2;159;163;169m9\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;247;247;248mÑ\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;47;59;73m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;86;121;142m=\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;138;196;228m$\x1b[0m\x1b[38;2;210;199;171mW\x1b[0m\x1b[38;2;66;78;88m~\x1b[0m\x1b[38;2;157;155;144m8\x1b[0m\x1b[38;2;204;196;173m$\x1b[0m\x1b[38;2;152;217;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;80;106;125m=\x1b[0m\x1b[38;2;105;144;172m8\x1b[0m\x1b[38;2;59;67;76m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;77;71;65m~\x1b[0m\x1b[38;2;97;91;86m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;13;13;13m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;193;198;202mW\x1b[0m\x1b[38;2;217;206;176mW\x1b[0m\x1b[38;2;207;197;169m$\x1b[0m\x1b[38;2;206;198;174mW\x1b[0m\x1b[38;2;80;92;108m+\x1b[0m\x1b[38;2;84;96;111m+\x1b[0m\x1b[38;2;41;54;71m-\x1b[0m\x1b[38;2;41;55;72m-\x1b[0m\x1b[38;2;44;57;74m-\x1b[0m\x1b[38;2;233;221;190m#\x1b[0m\x1b[38;2;58;72;88m~\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;49;58;72m-\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;59;72;87m~\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;50;62;77m-\x1b[0m\x1b[38;2;194;186;165m$\x1b[0m\x1b[38;2;97;104;108m=\x1b[0m\x1b[38;2;92;124;144m7\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;60;74;89m~\x1b[0m\x1b[38;2;113;158;188m8\x1b[0m\x1b[38;2;79;71;63m~\x1b[0m\x1b[38;2;183;188;194m$\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;30;30;30m,\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;241;242;243m@\x1b[0m\x1b[38;2;177;172;153m9\x1b[0m\x1b[38;2;207;197;170m$\x1b[0m\x1b[38;2;207;197;170m$\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;66;77;89m~\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;195;234m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;105;143;165m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;27;27;27m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;17;17;17m.\x1b[0m\x1b[38;2;230;230;230m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;60;72;87m~\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;195;234m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;67;79;93m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;250;250;250mÑ\x1b[0m\x1b[38;2;34;42;50m,\x1b[0m\x1b[38;2;53;65;78m~\x1b[0m\x1b[38;2;155;222;255m#\x1b[0m\x1b[38;2;234;235;237m@\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;104;104;104m=\x1b[0m\x1b[38;2;20;20;20m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;26;26;26m.\x1b[0m\x1b[38;2;21;21;21m.\x1b[0m\x1b[38;2;20;20;20m.\x1b[0m\x1b[38;2;20;20;20m.\x1b[0m\x1b[38;2;16;16;16m.\x1b[0m\x1b[38;2;34;34;34m,\x1b[0m\x1b[38;2;247;247;247mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;183;188;193m$\x1b[0m\x1b[38;2;140;199;237m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;195;233m$\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;39;48;57m-\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;178;184;190m$\x1b[0m\x1b[38;2;30;38;45m,\x1b[0m\x1b[38;2;31;39;46m,\x1b[0m\x1b[38;2;51;61;75m-\x1b[0m\x1b[38;2;50;61;76m-\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;27;27;27m.\x1b[0m\x1b[38;2;20;20;20m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;24;24;24m.\x1b[0m\x1b[38;2;100;100;100m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;69;81;95m~\x1b[0m\x1b[38;2;138;196;235m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;146;209;248mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;31;39;45m,\x1b[0m\x1b[38;2;30;38;44m,\x1b[0m\x1b[38;2;184;189;195m$\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;84;95;107m+\x1b[0m\x1b[38;2;82;93;106m+\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;224;224;224m#\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;26;26;26m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;32;32;32m,\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;65;78;93m~\x1b[0m\x1b[38;2;229;217;187m#\x1b[0m\x1b[38;2;76;100;120m+\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;196;234m$\x1b[0m\x1b[38;2;138;195;233m$\x1b[0m\x1b[38;2;145;205;244mW\x1b[0m\x1b[38;2;150;214;254mW\x1b[0m\x1b[38;2;150;215;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;150;214;255mW\x1b[0m\x1b[38;2;141;200;231m$\x1b[0m\x1b[38;2;49;61;76m-\x1b[0m\x1b[38;2;246;247;247mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;28;28;28m.\x1b[0m\x1b[38;2;21;21;21m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;128;129;125m7\x1b[0m\x1b[38;2;206;196;170m$\x1b[0m\x1b[38;2;224;212;184mW\x1b[0m\x1b[38;2;225;214;185mW\x1b[0m\x1b[38;2;169;166;152m9\x1b[0m\x1b[38;2;146;154;162m8\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;253;253;254mÑ\x1b[0m\x1b[38;2;136;144;153m8\x1b[0m\x1b[38;2;62;74;88m~\x1b[0m\x1b[38;2;49;61;75m-\x1b[0m\x1b[38;2;48;60;76m-\x1b[0m\x1b[38;2;53;66;81m~\x1b[0m\x1b[38;2;91;102;114m=\x1b[0m\x1b[38;2;190;195;200mW\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;203;203;203mW\x1b[0m\x1b[38;2;22;22;22m.\x1b[0m\x1b[38;2;33;33;33m,\x1b[0m\x1b[38;2;26;26;26m.\x1b[0m\x1b[38;2;20;20;20m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;21;21;21m.\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;78;90;103m+\x1b[0m\x1b[38;2;55;67;83m~\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;255;255;255mÑ\x1b[0m\x1b[38;2;105;105;105m=\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;30;30;30m,\x1b[0m\x1b[38;2;32;32;32m,\x1b[0m\x1b[38;2;32;32;32m,\x1b[0m\x1b[38;2;33;33;33m,\x1b[0m\x1b[38;2;31;31;31m,\x1b[0m\x1b[38;2;31;31;31m,\x1b[0m\x1b[38;2;30;30;30m,\x1b[0m\x1b[38;2;31;31;31m,\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;30;30;30m,\x1b[0m\x1b[38;2;28;28;28m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
	`\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;28;28;28m.\x1b[0m\x1b[38;2;30;30;30m,\x1b[0m\x1b[38;2;29;29;29m.\x1b[0m\x1b[38;2;24;24;24m.\x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\x1b[38;2;0;0;0m \x1b[0m\n`,
}
