package useCase

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/mocks"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {
	type args struct {
		order dto.RequestOrder
	}

	tests := []struct {
		name        string
		args        args
		wantOrder   *dto.OutputOrder
		wantError   bool
		mockGateway *mocks.OrderGatewayMock
	}{
		{
			name: "success",
			args: args{
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc123",
				},
			},
			wantOrder: &dto.OutputOrder{
				ID: 1,
				Client: dto.OutputClient{
					ID: 1,
				},
				VoucherID:        nil,
				Status:           "Done",
				VerificationCode: "abc123",
			},
			mockGateway: &mocks.OrderGatewayMock{
				WantOut: &dto.OutputOrder{
					ID: 1,
					Client: dto.OutputClient{
						ID: 1,
					},
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc123",
				},
				WantOutList: []dto.OutputOrder{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name: "error_verification_code",
			args: args{
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc1",
				},
			},
			wantOrder:   nil,
			mockGateway: &mocks.OrderGatewayMock{},
			wantError:   true,
		},
		{
			name: "error_invalid_status",
			args: args{
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "",
					VerificationCode: "abc123",
				},
			},
			wantOrder:   nil,
			mockGateway: &mocks.OrderGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		orderUseCase := NewOrderUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {

			o, err := orderUseCase.SaveOrder(tc.args.order)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				oWithoutVerifyCode := &dto.RequestOrder{
					ID:        o.ID,
					ClientID:  o.Client.ID,
					VoucherID: o.VoucherID,
					Status:    o.Status,
					CreatedAt: o.CreatedAt,
				}
				orderOutStr = ps.MarshalString(oWithoutVerifyCode)
			}

			if tc.wantOrder != nil {
				oWithoutVerifyCode := &dto.RequestOrder{
					ID:        tc.wantOrder.ID,
					ClientID:  tc.wantOrder.Client.ID,
					VoucherID: tc.wantOrder.VoucherID,
					Status:    tc.wantOrder.Status,
					CreatedAt: tc.wantOrder.CreatedAt,
				}
				wantOrderOutStr = ps.MarshalString(oWithoutVerifyCode)
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateOrderByID$
func Test_UpdateOrderByID(t *testing.T) {
	type args struct {
		id    int64
		order dto.RequestOrder
	}

	tests := []struct {
		name        string
		args        args
		wantOrder   *dto.OutputOrder
		wantError   bool
		mockGateway *mocks.OrderGatewayMock
	}{
		{
			name: "success",
			args: args{
				id: 1,
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc123",
				},
			},
			mockGateway: &mocks.OrderGatewayMock{
				WantOut: &dto.OutputOrder{
					ID:               1,
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc123",
				},
				WantOutList: []dto.OutputOrder{},
				WantErr:     nil,
			},
			wantOrder: &dto.OutputOrder{
				ID:               1,
				VoucherID:        nil,
				Status:           "Done",
				VerificationCode: "abc123",
			},
			wantError: false,
		},
		{
			name: "error_invalid_status",
			args: args{
				id: 1,
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "",
					VerificationCode: "abc1",
				},
			},
			wantOrder:   nil,
			mockGateway: &mocks.OrderGatewayMock{},
			wantError:   true,
		},
		{
			name: "error_invalid_id",
			args: args{
				order: dto.RequestOrder{
					ID:               1,
					ClientID:         1,
					VoucherID:        nil,
					Status:           "",
					VerificationCode: "abc1",
				},
			},
			wantOrder:   nil,
			mockGateway: &mocks.OrderGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		orderUseCase := NewOrderUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			o, err := orderUseCase.UpdateOrderByID(tc.args.id, tc.args.order)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				oWithoutVerifyCode := &dto.RequestOrder{
					ID:        o.ID,
					ClientID:  o.Client.ID,
					VoucherID: o.VoucherID,
					Status:    o.Status,
					CreatedAt: o.CreatedAt,
				}
				orderOutStr = ps.MarshalString(oWithoutVerifyCode)
			}

			if tc.wantOrder != nil {
				oWithoutVerifyCode := &dto.RequestOrder{
					ID:        tc.wantOrder.ID,
					ClientID:  tc.wantOrder.Client.ID,
					VoucherID: tc.wantOrder.VoucherID,
					Status:    tc.wantOrder.Status,
					CreatedAt: tc.wantOrder.CreatedAt,
				}
				wantOrderOutStr = ps.MarshalString(oWithoutVerifyCode)
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetOrderByID$
func Test_GetOrderByID(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		wantOrder   *dto.OutputOrder
		wantError   bool
		mockGateway *mocks.OrderGatewayMock
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantOrder: &dto.OutputOrder{
				ID:               1,
				VoucherID:        nil,
				Status:           "Done",
				VerificationCode: "abc123",
			},
			mockGateway: &mocks.OrderGatewayMock{
				WantOut: &dto.OutputOrder{
					ID:               1,
					VoucherID:        nil,
					Status:           "Done",
					VerificationCode: "abc123",
				},
				WantOutList: []dto.OutputOrder{},
				WantErr:     nil,
			},
			wantError: false,
		},
		{
			name:        "error_invalid_id",
			wantOrder:   nil,
			mockGateway: &mocks.OrderGatewayMock{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		orderUseCase := NewOrderUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			_, err := orderUseCase.GetOrderByID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}
