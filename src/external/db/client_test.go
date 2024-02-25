package db

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/src/base/mocks"
	"fiap-hf-src/src/core/entity"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -cover -run ^Test_GetClientByID$
func Test_GetClientByID(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		ClientID int64
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Client
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{
				ClientID: 0,
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name: "connection_error",
			args: args{
				ClientID: 0,
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name: "err_query",
			args: args{
				ClientID: 0,
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errQuery"),
			},
			isWantError: true,
		},
		{
			name: "prepare_stmt_error",
			args: args{
				ClientID: 0,
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name: "error_scan",
			args: args{},
			ctx:  nil,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
				CreatedAt: com.CreatedAt{
					Value: time.Time{},
				},
			},
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
			repo := NewClientDB(tc.ctx, tc.mockDB)

			out, err := repo.GetClientByID(tc.args.ClientID)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		ClientCPF string
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Client
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{
				ClientCPF: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name: "connection_error",
			args: args{
				ClientCPF: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name: "err_query",
			args: args{
				ClientCPF: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errQuery"),
			},
			isWantError: true,
		},
		{
			name: "prepare_stmt_error",
			args: args{
				ClientCPF: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name: "error_scan",
			args: args{},
			ctx:  nil,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
				CreatedAt: com.CreatedAt{
					Value: time.Time{},
				},
			},
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
			repo := NewClientDB(tc.ctx, tc.mockDB)

			out, err := repo.GetClientByCPF(tc.args.ClientCPF)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {
	printArt()
	ctx := context.Background()

	type args struct {
		client entity.Client
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Client
		mockDB      *mocks.MockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{
				client: entity.Client{
					ID:   0,
					Name: "",
					CPF: com.Cpf{
						Value: "",
					},
					Email: "",
					CreatedAt: com.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name: "connection_error",
			args: args{},
			ctx:  ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
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
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mocks.MockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errScan"),
			},

			isWantError: true,
		},
		{
			name: "error_scan_stmt",
			args: args{},
			ctx:  nil,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: com.Cpf{
					Value: "",
				},
				Email: "",
				CreatedAt: com.CreatedAt{
					Value: time.Time{},
				},
			},
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
			repo := NewClientDB(tc.ctx, tc.mockDB)

			out, err := repo.SaveClient(tc.args.client)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.WantOutput)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.WantOutput), ps.MarshalString(out))
			}

		})
	}
}
