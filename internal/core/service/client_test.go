package service

import (
	"fiap-hf-src/internal/core/entity"
	"fiap-hf-src/internal/core/entity/common"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {

	serviceClient := NewClientService(nil)

	type args struct {
		client entity.Client
	}

	tests := []struct {
		name       string
		args       args
		wantClient *entity.Client
		wantError  bool
	}{
		{
			name: "success",
			args: args{
				client: entity.Client{
					Name: "Doc Emmet Brown",
					CPF: common.Cpf{
						Value: "12345",
					},
					Email: "doc@delorean.com",
				},
			},
			wantClient: &entity.Client{
				Name: "Doc Emmet Brown",
				CPF: common.Cpf{
					Value: "12345",
				},
				Email: "doc@delorean.com",
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				client: entity.Client{
					Name: "Doc Emmet Brown",
					CPF: common.Cpf{
						Value: "",
					},
					Email: "doc@delorean.com",
				},
			},
			wantClient: nil,
			wantError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			c, err := serviceClient.SaveClient(tc.args.client)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			clientOutStr, wantClientOutStr := "nil", "nil"
			if c != nil {
				clientOutStr = c.MarshalString()
			}

			if tc.wantClient != nil {
				wantClientOutStr = tc.wantClient.MarshalString()
			}

			if clientOutStr != wantClientOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantClientOutStr, clientOutStr)
			}
		})
	}

}

// go test -v -count=1 -failfast -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	newClient := entity.Client{}

	clientService := NewClientService(&newClient)

	type args struct {
		cpf string
	}

	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "success",
			args: args{
				cpf: "some",
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				cpf: "",
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			gotErr := clientService.GetClientByCPF(tc.args.cpf)

			if (!tc.wantError) && gotErr != nil {
				t.Fatalf("error: %v", gotErr)
			}

		})
	}

}

// go test -v -count=1 -failfast -run ^Test_GetClientByID$
func Test_GetClientByID(t *testing.T) {
	newClient := entity.Client{}

	clientService := NewClientService(&newClient)

	type args struct {
		id int64
	}

	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantError: false,
		},
		{
			name: "error",
			args: args{
				id: 0,
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			gotErr := clientService.GetClientByID(tc.args.id)

			if (!tc.wantError) && gotErr != nil {
				t.Fatalf("error: %v", gotErr)
			}

		})
	}
}
