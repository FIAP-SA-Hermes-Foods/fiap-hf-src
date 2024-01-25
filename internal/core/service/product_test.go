package service

import (
	"fiap-hf-src/internal/core/entity"
	"fiap-hf-src/internal/core/entity/common"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -run ^Test_SaveProduct$
func Test_SaveProduct(t *testing.T) {
	serviceProduct := NewProductService(nil)

	type args struct {
		product entity.Product
	}

	tests := []struct {
		name        string
		args        args
		wantProduct *entity.Product
		wantError   bool
	}{
		{
			name: "success_english_category",
			args: args{
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "drink",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: &entity.Product{
				ID:   1,
				Name: "",
				Category: common.Category{
					Value: "DRINK",
				},
				Image:       "",
				Description: "",
				Price:       0.0,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				DeactivatedAt: common.DeactivatedAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "success_portuguese_category",
			args: args{
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "acompanhamento",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: &entity.Product{
				ID:   1,
				Name: "",
				Category: common.Category{
					Value: "COMPLEMENT",
				},
				Image:       "",
				Description: "",
				Price:       0.0,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				DeactivatedAt: common.DeactivatedAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_invalid_category",
			args: args{
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "acompanhament",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			p, err := serviceProduct.SaveProduct(tc.args.product)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			pStr, wantPStr := "nil", "nil"

			if p != nil {
				pStr = p.MarshalString()
			}

			if tc.wantProduct != nil {
				wantPStr = tc.wantProduct.MarshalString()
			}

			if pStr != wantPStr {
				t.Fatalf("want: %v\ngot: %v", wantPStr, pStr)

			}
		})
	}

}

// go test -v -count=1 -failfast -run ^Test_UpdateProductByID$
func Test_UpdateProductByID(t *testing.T) {
	serviceProduct := NewProductService(nil)

	type args struct {
		id      int64
		product entity.Product
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		wantProduct *entity.Product
	}{
		{
			name: "success_english_category",
			args: args{
				id: 1,
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "drink",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: &entity.Product{
				ID:   1,
				Name: "",
				Category: common.Category{
					Value: "DRINK",
				},
				Image:       "",
				Description: "",
				Price:       0.0,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				DeactivatedAt: common.DeactivatedAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "success_portuguese_category",
			args: args{
				id: 1,
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "acompanhamento",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: &entity.Product{
				ID:   1,
				Name: "",
				Category: common.Category{
					Value: "COMPLEMENT",
				},
				Image:       "",
				Description: "",
				Price:       0.0,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				DeactivatedAt: common.DeactivatedAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_invalid_category",
			args: args{
				id: 1,
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "acompanhament",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_id",
			args: args{
				product: entity.Product{
					ID:   1,
					Name: "",
					Category: common.Category{
						Value: "acompanhament",
					},
					Image:       "",
					Description: "",
					Price:       0.0,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					DeactivatedAt: common.DeactivatedAt{
						Value: &time.Time{},
					},
				},
			},
			wantProduct: nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			p, err := serviceProduct.UpdateProductByID(tc.args.id, tc.args.product)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			pStr, wantPStr := "nil", "nil"

			if p != nil {
				pStr = p.MarshalString()
			}

			if tc.wantProduct != nil {
				wantPStr = tc.wantProduct.MarshalString()
			}

			if pStr != wantPStr {
				t.Fatalf("want: %v\ngot: %v", wantPStr, pStr)

			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByID$
func Test_GetProductByID(t *testing.T) {
	serviceProduct := NewProductService(nil)

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
			name: "error_invalid_id",
			args: args{
				id: 0,
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			err := serviceProduct.GetProductByID(tc.args.id)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetProductByCategory$
func Test_GetProductByCategory(t *testing.T) {
	serviceProduct := NewProductService(nil)

	type args struct {
		category string
	}

	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "success",
			args: args{
				category: "meal",
			},
			wantError: false,
		},
		{
			name: "error_empty_category",
			args: args{
				category: "",
			},
			wantError: true,
		},
		{
			name: "error_invalid_category",
			args: args{
				category: "drinke",
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			err := serviceProduct.GetProductByCategory(tc.args.category)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_DeleteProductByI$
func Test_DeleteProductByI(t *testing.T) {
	p := entity.Product{}
	serviceProduct := NewProductService(&p)

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
			name: "error_invalid_id",
			args: args{
				id: 0,
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			err := serviceProduct.DeleteProductByID(tc.args.id)
			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}
