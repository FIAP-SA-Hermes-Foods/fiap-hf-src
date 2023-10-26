package api_mercadoPago

import (
	"context"
	"errors"
	"hermes-foods/internal/core/domain/entity"
	"hermes-foods/internal/core/domain/valueObject"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

// go test -v -count=1 -cover -failfast -run ^Test_Integration$
func Test_Integration(t *testing.T) {

	ctx := context.Background()
	// ctx = context.WithValue(ctx, )

	in := entity.InputPaymentAPI{
		Price: 0.0,
		Client: entity.Client{
			Name: "",
			CPF: valueObject.Cpf{
				Value: "",
			},
			Email: "",
		},
	}

	headers := map[string]string{"Content-Type": "application/json"}

	api := NewMercadoPagoAPI("http://127.0.0.1:8081/mercado_pago_api", headers, time.Second*5)

	out, err := api.DoPayment(ctx, in)

	if out != nil {

		log.Printf("Output -> %s\nErr -> %v", out.MarshalString(), err)
		return
	}

	log.Printf("Err -> %v", err)

}

// go test -v -count=1 -cover -failfast -run ^Test_DoPayment$
func Test_DoPayment(t *testing.T) {

	ctx := context.Background()

	du, _ := time.ParseDuration("5000ms")

	type args struct {
		ctx   context.Context
		input entity.InputPaymentAPI
	}

	returnPaidAPI := mockOutputDoPayment(successDoPaymentOutputMock)
	returnFailAPI := mockOutputDoPayment(errorDoPaymentOutputMock)

	tests := []struct {
		name       string
		args       args
		url        string
		headers    map[string]string
		timeout    time.Duration
		mock       *httpMock
		wantOutput *entity.OutputPaymentAPI
		isWantErr  bool
	}{
		{
			name: "success_mercado_pago_output",
			args: args{
				ctx:   ctx,
				input: mockInputDoPayment(successInput),
			},
			url: "http://localhost:8000",
			headers: map[string]string{
				"": "",
			},
			timeout:    du,
			wantOutput: &returnPaidAPI,
			mock: &httpMock{
				WantOut: &http.Response{
					Status:     "OK",
					StatusCode: 200,
					Header: map[string][]string{
						"Content-type": {
							"application/json",
						},
					},
					Body: io.NopCloser(strings.NewReader(successResponseAPIMock)),
				},
				WantErr: nil,
			},
			isWantErr: false,
		},
		{
			name: "error_returned_from_request",
			args: args{
				ctx:   ctx,
				input: mockInputDoPayment(successInput),
			},
			url: "http://localhost:8000",
			headers: map[string]string{
				"": "",
			},
			timeout:    du,
			wantOutput: &returnPaidAPI,
			mock: &httpMock{
				WantOut: &http.Response{
					Status:     "OK",
					StatusCode: 200,
					Header: map[string][]string{
						"Content-type": {
							"application/json",
						},
					},
					Body: io.NopCloser(strings.NewReader(successDoPaymentOutputMock)),
				},
				WantErr: errors.New("Great Scott! The Delorean is without fuel!"),
			},
			isWantErr: true,
		},
		{
			name: "error_request_body",
			args: args{
				ctx:   ctx,
				input: mockInputDoPayment(successInput),
			},
			url: "http://localhost:8000",
			headers: map[string]string{
				"": "",
			},
			timeout:    du,
			wantOutput: &returnPaidAPI,
			mock: &httpMock{
				WantOut: &http.Response{
					Status:     "OK",
					StatusCode: 200,
					Header: map[string][]string{
						"Content-type": {
							"application/json",
						},
					},
					Body: io.NopCloser(strings.NewReader("<==>")),
				},
				WantErr: nil,
			},
			isWantErr: true,
		},
		{
			name: "error_returned_from_api",
			args: args{
				ctx:   ctx,
				input: mockInputDoPayment(successInput),
			},
			url: "http://localhost:8000",
			headers: map[string]string{
				"": "",
			},
			timeout:    du,
			wantOutput: &returnFailAPI,
			mock: &httpMock{
				WantOut: &http.Response{
					Status:     "Internal Server Error",
					StatusCode: 500,
					Header: map[string][]string{
						"Content-type": {
							"application/json",
						},
					},
					Body: io.NopCloser(strings.NewReader(errorDoPaymentOutputMock)),
				},
				WantErr: nil,
			},
			isWantErr: true,
		},
	}

	for _, tc := range tests {
		httpClient = tc.mock
		t.Run(tc.name, func(t *testing.T) {

			api := NewMercadoPagoAPI(tc.url, tc.headers, tc.timeout)

			resp, err := api.DoPayment(tc.args.ctx, tc.args.input)

			if (!tc.isWantErr) && err != nil {
				t.Errorf("was suppose to not have null error here and %v got", err)
				return
			}

			if resp != nil && resp.MarshalString() != tc.wantOutput.MarshalString() {
				t.Errorf("was suppose to have: %s\nand got:%s\n", tc.wantOutput.MarshalString(), resp.MarshalString())
				return
			}
		})
	}
}

func Test_mockInputDoPayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		in := mockInputDoPayment(successInput)
		if in == (entity.InputPaymentAPI{}) {
			t.Errorf("fail to convert")
		}
	})

	t.Run("error", func(t *testing.T) {
		in := mockInputDoPayment("<==>")
		if in != (entity.InputPaymentAPI{}) {
			t.Errorf("stranger things happen here")
		}
	})
}

func Test_mockOutputDoPayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		out := mockOutputDoPayment(successDoPaymentOutputMock)
		if out == (entity.OutputPaymentAPI{}) {
			t.Errorf("fail to convert")
		}
	})

	t.Run("error", func(t *testing.T) {
		out := mockOutputDoPayment("<==>")
		if out != (entity.OutputPaymentAPI{}) {
			t.Errorf("stranger things happen here")
		}
	})
}
