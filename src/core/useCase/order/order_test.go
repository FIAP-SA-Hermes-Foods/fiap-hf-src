package service

import (
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/core/entity/common"
	"fiap-hf-src/src/operation/presenter"
	"testing"
	"time"
)

// go test -v -count=1 -failfast -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {
	orderService := NewOrderService(nil)

	type args struct {
		order entity.Order
	}

	tests := []struct {
		name      string
		args      args
		wantOrder *entity.Order
		wantError bool
	}{
		{
			name: "success",
			args: args{
				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "Done",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc123",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: &entity.Order{
				ID:        1,
				ClientID:  1,
				VoucherID: nil,
				Items:     []entity.OrderItems{},
				Status: common.Status{
					Value: "Done",
				},
				VerificationCode: common.VerificationCode{
					Value: "abc123",
				},
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_verification_code",
			args: args{
				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "Done",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc1",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: &entity.Order{
				ID:        1,
				ClientID:  1,
				VoucherID: nil,
				Items:     []entity.OrderItems{},
				Status: common.Status{
					Value: "Done",
				},
				VerificationCode: common.VerificationCode{
					Value: "abc123",
				},
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
			},
			wantError: true,
		},
		{
			name: "error_invalid_status",
			args: args{
				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc1",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			o, err := orderService.SaveOrder(tc.args.order)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				oWithoutVerifyCode := &entity.Order{
					ID:        o.ID,
					ClientID:  o.ClientID,
					VoucherID: o.VoucherID,
					Items:     o.Items,
					Status:    o.Status,
					CreatedAt: o.CreatedAt,
				}
				orderOutStr = presenter.MarshalString(oWithoutVerifyCode)
			}

			if tc.wantOrder != nil {
				oWithoutVerifyCode := &entity.Order{
					ID:        tc.wantOrder.ID,
					ClientID:  tc.wantOrder.ClientID,
					VoucherID: tc.wantOrder.VoucherID,
					Items:     tc.wantOrder.Items,
					Status:    tc.wantOrder.Status,
					CreatedAt: tc.wantOrder.CreatedAt,
				}
				wantOrderOutStr = presenter.MarshalString(oWithoutVerifyCode)
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateOrderByID$
func Test_UpdateOrderByID(t *testing.T) {
	newO := entity.Order{}
	orderService := NewOrderService(&newO)

	type args struct {
		id    int64
		order entity.Order
	}

	tests := []struct {
		name      string
		args      args
		wantOrder *entity.Order
		wantError bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "Done",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc123",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: &entity.Order{
				ID:        1,
				ClientID:  1,
				VoucherID: nil,
				Items:     []entity.OrderItems{},
				Status: common.Status{
					Value: "Done",
				},
				VerificationCode: common.VerificationCode{
					Value: "abc123",
				},
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_invalid_status",
			args: args{
				id: 1,
				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc1",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: nil,
			wantError: true,
		},
		{
			name: "error_invalid_id",
			args: args{

				order: entity.Order{
					ID:        1,
					ClientID:  1,
					VoucherID: nil,
					Items:     []entity.OrderItems{},
					Status: common.Status{
						Value: "",
					},
					VerificationCode: common.VerificationCode{
						Value: "abc1",
					},
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
				},
			},
			wantOrder: nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			o, err := orderService.UpdateOrderByID(tc.args.id, tc.args.order)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			orderOutStr, wantOrderOutStr := "nil", "nil"

			if o != nil {
				oWithoutVerifyCode := &entity.Order{
					ID:        o.ID,
					ClientID:  o.ClientID,
					VoucherID: o.VoucherID,
					Items:     o.Items,
					Status:    o.Status,
					CreatedAt: o.CreatedAt,
				}
				orderOutStr = presenter.MarshalString(oWithoutVerifyCode)
			}

			if tc.wantOrder != nil {
				oWithoutVerifyCode := &entity.Order{
					ID:        tc.wantOrder.ID,
					ClientID:  tc.wantOrder.ClientID,
					VoucherID: tc.wantOrder.VoucherID,
					Items:     tc.wantOrder.Items,
					Status:    tc.wantOrder.Status,
					CreatedAt: tc.wantOrder.CreatedAt,
				}
				wantOrderOutStr = presenter.MarshalString(oWithoutVerifyCode)
			}

			if orderOutStr != wantOrderOutStr {
				t.Fatalf("was suppose to have %s and %s got", wantOrderOutStr, orderOutStr)
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_GetOrderByID$
func Test_GetOrderByID(t *testing.T) {
	newO := entity.Order{}
	orderService := NewOrderService(&newO)

	type args struct {
		id int64
	}

	tests := []struct {
		name      string
		args      args
		wantOrder *entity.Order
		wantError bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantOrder: &entity.Order{
				ID:        1,
				ClientID:  1,
				VoucherID: nil,
				Items:     []entity.OrderItems{},
				Status: common.Status{
					Value: "Done",
				},
				VerificationCode: common.VerificationCode{
					Value: "abc123",
				},
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
			},
			wantError: false,
		},
		{
			name:      "error_invalid_id",
			wantOrder: nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {

			err := orderService.GetOrderByID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

		})
	}

}
