package service

import (
	"fiap-hf-src/internal/core/entity"
	"fiap-hf-src/internal/core/entity/common"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -run ^Test_GetOrderProductByOrderID$
func Test_GetOrderProductByOrderID(t *testing.T) {
	newO := entity.OrderProduct{}
	orderService := NewOrderProductService(&newO)

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
			name:      "error_invalid_id",
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			err := orderService.GetOrderProductByOrderID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

		})
	}

}

// go test -v -count=1 -failfast -run ^Test_SaveOrderProduct$
func Test_SaveOrderProduct(t *testing.T) {
	orderProductService := NewOrderProductService(nil)

	type args struct {
		orderProduct entity.OrderProduct
	}

	tests := []struct {
		name      string
		args      args
		wantOrder *entity.OrderProduct
		wantError bool
	}{
		{
			name: "success",
			args: args{
				orderProduct: entity.OrderProduct{
					ID:         1,
					Quantity:   1,
					TotalPrice: 20.0,
					Discount:   10.0,
					OrderID:    1,
					ProductID:  nil,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: &entity.OrderProduct{
				ID:         1,
				Quantity:   1,
				TotalPrice: 20.0,
				Discount:   10.0,
				OrderID:    1,
				ProductID:  nil,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
			},
			wantError: false,
		},

		{
			name:      "error_invalid_id",
			args:      args{},
			wantOrder: nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			o, err := orderProductService.SaveOrderProduct(tc.args.orderProduct)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				orderOutStr = o.MarshalString()
			}

			if tc.wantOrder != nil {
				wantOrderOutStr = tc.wantOrder.MarshalString()
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}
