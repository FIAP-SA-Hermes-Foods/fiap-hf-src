package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/mocks"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {

	type args struct {
		client dto.RequestClient
	}

	tests := []struct {
		name        string
		args        args
		wantClient  *dto.OutputClient
		wantError   bool
		mockGateway *mocks.ClientGatewayMock
	}{
		{
			name: "success",
			args: args{
				client: dto.RequestClient{
					Name:  "Doc Emmet Brown",
					CPF:   "12345",
					Email: "doc@delorean.com",
				},
			},
			wantClient: &dto.OutputClient{
				Name:  "Doc Emmet Brown",
				CPF:   "12345",
				Email: "doc@delorean.com",
			},
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: &dto.OutputClient{
					Name:  "Doc Emmet Brown",
					CPF:   "12345",
					Email: "doc@delorean.com",
				},
				WantErr: nil,
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				client: dto.RequestClient{
					Name:  "Doc Emmet Brown",
					CPF:   "",
					Email: "doc@delorean.com",
				},
			},
			wantClient: nil,
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: nil,
				WantErr: errors.New("errSaveClient"),
			},
			wantError: true,
		},
	}

	for _, tc := range tests {

		useCaseClient := NewClientUseCase(tc.mockGateway)

		t.Run(tc.name, func(*testing.T) {

			c, err := useCaseClient.SaveClient(tc.args.client)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			clientOutStr, wantClientOutStr := "nil", "nil"
			if c != nil {
				clientOutStr = ps.MarshalString(c)
			}

			if tc.wantClient != nil {
				wantClientOutStr = ps.MarshalString(tc.wantClient)
			}

			if clientOutStr != wantClientOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantClientOutStr, clientOutStr)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	type args struct {
		cpf string
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		mockGateway *mocks.ClientGatewayMock
	}{
		{
			name: "success",
			args: args{
				cpf: "some",
			},
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: &dto.OutputClient{
					CPF: "some",
				},
				WantErr: nil,
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				cpf: "",
			},
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: nil,
				WantErr: errors.New("errGetClientByCPF"),
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		useCaseClient := NewClientUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {

			_, gotErr := useCaseClient.GetClientByCPF(tc.args.cpf)

			if (!tc.wantError) && gotErr != nil {
				t.Fatalf("error: %v", gotErr)
			}

		})
	}

}

// go test -v -count=1 -failfast -run ^Test_GetClientByID$
func Test_GetClientByID(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		mockGateway *mocks.ClientGatewayMock
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: &dto.OutputClient{
					ID: 1,
				},
				WantErr: nil,
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				id: 0,
			},
			mockGateway: &mocks.ClientGatewayMock{
				WantOut: nil,
				WantErr: errors.New("errGetClientByID"),
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		useCaseClient := NewClientUseCase(tc.mockGateway)

		t.Run(tc.name, func(*testing.T) {

			_, gotErr := useCaseClient.GetClientByID(tc.args.id)

			if (!tc.wantError) && gotErr != nil {
				t.Fatalf("error: %v", gotErr)
			}

		})
	}
}
