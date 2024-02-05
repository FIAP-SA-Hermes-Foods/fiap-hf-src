package useCase

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/mocks"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {

	type args struct {
		product dto.RequestProduct
	}

	tests := []struct {
		name        string
		args        args
		wantProduct *dto.OutputProduct
		mockGateway *mocks.ProductGatewayMock
		wantError   bool
	}{
		{
			name: "success_english_category",
			args: args{
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "drink",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: &dto.OutputProduct{
				ID:          1,
				Name:        "",
				Category:    "DRINK",
				Image:       "",
				Description: "",
				Price:       0.0,
			},
			mockGateway: &mocks.ProductGatewayMock{
				WantOut: &dto.OutputProduct{
					ID:          1,
					Name:        "",
					Category:    "DRINK",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
				WantOutList: []dto.OutputProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "success_portuguese_category",
			args: args{
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "acompanhamento",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: &dto.OutputProduct{
				ID:          1,
				Name:        "",
				Category:    "COMPLEMENT",
				Image:       "",
				Description: "",
				Price:       0.0,
			},
			mockGateway: &mocks.ProductGatewayMock{
				WantOut: &dto.OutputProduct{
					ID:          1,
					Name:        "",
					Category:    "COMPLEMENT",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
				WantOutList: []dto.OutputProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "error_invalid_category",
			args: args{
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "acompanhament",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: nil,
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceProduct := NewProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {

			p, err := serviceProduct.SaveProduct(tc.args.product)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			pStr, wantPStr := "nil", "nil"

			if p != nil {
				pStr = ps.MarshalString(p)
			}

			if tc.wantProduct != nil {
				wantPStr = ps.MarshalString(tc.wantProduct)
			}

			if pStr != wantPStr {
				t.Fatalf("want: %v\ngot: %v", wantPStr, pStr)

			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateProductByID$
func Test_UpdateProductByID(t *testing.T) {

	type args struct {
		id      int64
		product dto.RequestProduct
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		mockGateway *mocks.ProductGatewayMock
		wantProduct *dto.OutputProduct
	}{
		{
			name: "success_english_category",
			args: args{
				id: 1,
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "drink",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: &dto.OutputProduct{
				ID:          1,
				Name:        "",
				Category:    "DRINK",
				Image:       "",
				Description: "",
				Price:       0.0,
			},
			mockGateway: &mocks.ProductGatewayMock{
				WantOut: &dto.OutputProduct{
					ID:          1,
					Name:        "",
					Category:    "DRINK",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
				WantOutList: []dto.OutputProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "success_portuguese_category",
			args: args{
				id: 1,
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "acompanhamento",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: &dto.OutputProduct{
				ID:          1,
				Name:        "",
				Category:    "COMPLEMENT",
				Image:       "",
				Description: "",
				Price:       0.0,
			},
			mockGateway: &mocks.ProductGatewayMock{
				WantOut: &dto.OutputProduct{
					ID:          1,
					Name:        "",
					Category:    "COMPLEMENT",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
				WantOutList: []dto.OutputProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "error_invalid_category",
			args: args{
				id: 1,
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "acompanhament",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: nil,
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
		{
			name: "error_invalid_id",
			args: args{
				product: dto.RequestProduct{
					ID:          1,
					Name:        "",
					Category:    "acompanhament",
					Image:       "",
					Description: "",
					Price:       0.0,
				},
			},
			wantProduct: nil,
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceProduct := NewProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			p, err := serviceProduct.UpdateProductByID(tc.args.id, tc.args.product)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			pStr, wantPStr := "nil", "nil"

			if p != nil {
				pStr = ps.MarshalString(p)
			}

			if tc.wantProduct != nil {
				wantPStr = ps.MarshalString(tc.wantProduct)
			}

			if pStr != wantPStr {
				t.Fatalf("want: %v\ngot: %v", wantPStr, pStr)

			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByID$
func Test_GetProductByID(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		mockGateway *mocks.ProductGatewayMock
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   false,
		},
		{
			name: "error_invalid_id",
			args: args{
				id: 0,
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceProduct := NewProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			_, err := serviceProduct.GetProductByID(tc.args.id)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {

	type args struct {
		category string
	}

	tests := []struct {
		name        string
		args        args
		mockGateway *mocks.ProductGatewayMock
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				category: "meal",
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   false,
		},
		{
			name: "error_empty_category",
			args: args{
				category: "",
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
		{
			name: "error_invalid_category",
			args: args{
				category: "drinke",
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceProduct := NewProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			_, err := serviceProduct.GetProductByCategory(tc.args.category)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_DeleteProductByI$
func Test_DeleteProductByI(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		mockGateway *mocks.ProductGatewayMock
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   false,
		},
		{
			name: "error_invalid_id",
			args: args{
				id: 0,
			},
			mockGateway: &mocks.ProductGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceProduct := NewProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			err := serviceProduct.DeleteProductByID(tc.args.id)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}
