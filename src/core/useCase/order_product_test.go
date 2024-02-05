package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/mocks"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_GetOrderProductByOrderID$
func Test_GetOrderProductByOrderID(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		mockGateway *mocks.OrderProductGatewayMock
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut:     &dto.OutputOrderProduct{},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "error_invalid_id",
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut:     &dto.OutputOrderProduct{},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     errors.New("errGetOrderProductByOrderID"),
			},
			wantError: true,
		},
		{
			name: "error_gateway",
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut:     &dto.OutputOrderProduct{},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     errors.New("errGetOrderProductByOrderID"),
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		orderUseCase := NewOrderProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {

			_, err := orderUseCase.GetAllOrderProductByOrderID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

		})
	}

}

// go test -v -count=1 -failfast -run ^Test_SaveOrderProduct$
func Test_SaveOrderProduct(t *testing.T) {

	type args struct {
		orderProduct dto.RequestOrderProduct
	}

	tests := []struct {
		name        string
		args        args
		wantOrder   *dto.OutputOrderProduct
		mockGateway *mocks.OrderProductGatewayMock
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				orderProduct: dto.RequestOrderProduct{
					ID:         1,
					Quantity:   1,
					TotalPrice: 20.0,
					Discount:   10.0,
					OrderID:    1,
					ProductID:  nil,
				},
			},
			wantOrder: &dto.OutputOrderProduct{
				ID:         1,
				Quantity:   1,
				TotalPrice: 20.0,
				Discount:   10.0,
				OrderID:    1,
				ProductID:  nil,
			},
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut: &dto.OutputOrderProduct{
					ID:         1,
					Quantity:   1,
					TotalPrice: 20.0,
					Discount:   10.0,
					OrderID:    1,
					ProductID:  nil,
				},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     nil,
			},
			wantError: false,
		},

		{
			name:      "error_invalid_id",
			args:      args{},
			wantOrder: nil,
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut:     &dto.OutputOrderProduct{},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     errors.New("errSaveOrderProduct"),
			},
			wantError: true,
		},
		{
			name: "error_gateway",
			args: args{
				orderProduct: dto.RequestOrderProduct{
					ID:         1,
					Quantity:   1,
					TotalPrice: 20.0,
					Discount:   10.0,
					OrderID:    1,
					ProductID:  nil,
				},
			},
			wantOrder: nil,
			mockGateway: &mocks.OrderProductGatewayMock{
				WantOut:     &dto.OutputOrderProduct{},
				WantOutList: []dto.OutputOrderProduct{},
				WantErr:     errors.New("errSaveOrderProduct"),
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		orderProductUseCase := NewOrderProductUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {

			o, err := orderProductUseCase.SaveOrderProduct(tc.args.orderProduct)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				orderOutStr = ps.MarshalString(o)
			}

			if tc.wantOrder != nil {
				wantOrderOutStr = ps.MarshalString(tc.wantOrder)
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}
