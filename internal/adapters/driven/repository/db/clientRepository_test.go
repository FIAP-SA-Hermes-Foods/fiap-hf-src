package db

import (
	"context"
	"database/sql"
	"errors"
	"hermes-foods/internal/core/domain/entity"
	"hermes-foods/internal/core/domain/valueObject"
	"testing"
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
				WantErr:    errors.New("errPrepareStmt"),
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewClientRepository(tc.ctx, tc.mockDB)

			out, err := repo.GetClientByID(tc.args.ClientID)

			if (!tc.isWantError) && err != nil {
				t.Errorf("was not suppose to have an error here and %v got", err)
			}

			if out != nil && (out.MarshalString() != tc.WantOutput.MarshalString()) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", tc.WantOutput.MarshalString(), out.MarshalString())
			}

		})
	}
}
