package client

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -cover -run ^Test_GetClientByID$
func Test_GetClientByID(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ClientID string
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Client
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{
				ClientID: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},

			isWantError: false,
		},
		{
			name: "connection_error",
			args: args{
				ClientID: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errConnect"),
			},

			isWantError: true,
		},
		{
			name: "err_query",
			args: args{
				ClientID: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    errors.New("errQuery"),
			},
			isWantError: true,
		},
		{
			name: "prepare_stmt_error",
			args: args{
				ClientID: "",
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
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
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
			},
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
			repo := NewClientRepository(tc.ctx, tc.mockDB)

			out, err := repo.GetClientByCPF(tc.args.ClientID)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}
		})
	}
}

// go test -v -count=1 -failfast -cover -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {
	ctx := context.Background()

	type args struct {
		client entity.Client
	}

	tests := []struct {
		name        string
		args        args
		ctx         context.Context
		WantOutput  *entity.Client
		mockDB      *mockDb
		isWantError bool
	}{
		{
			name: "success",
			args: args{
				client: entity.Client{
					ID:   0,
					Name: "",
					CPF: valueObject.Cpf{
						Value: "",
					},
					Email: "",
					CreatedAt: valueObject.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			ctx: ctx,
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
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
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
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
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
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
			WantOutput: &entity.Client{
				ID:   0,
				Name: "",
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
			},
			mockDB: &mockDb{
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
				CPF: valueObject.Cpf{
					Value: "",
				},
				Email: "",
				CreatedAt: valueObject.CreatedAt{
					Value: time.Time{},
				},
			},
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
			repo := NewClientRepository(tc.ctx, tc.mockDB)

			out, err := repo.SaveClient(tc.args.client)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}

		})
	}
}
