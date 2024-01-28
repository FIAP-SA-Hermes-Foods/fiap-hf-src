package service

import (
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/core/entity/common"
	"testing"
	"time"
)

// go test -v -failfast -cover -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {
	serviceVoucher := NewVoucherService(nil)

	type args struct {
		voucher entity.Voucher
	}

	tests := []struct {
		name        string
		args        args
		wantVoucher *entity.Voucher
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				voucher: entity.Voucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: &entity.Voucher{
				ID:         1,
				Code:       "SUCCESS10",
				Porcentage: 10,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: common.ExpiresAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_invalid_code",
			args: args{
				voucher: entity.Voucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_upper",
			args: args{
				voucher: entity.Voucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: 1000,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_lower",
			args: args{
				voucher: entity.Voucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: -1,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_expires_at_value",
			args: args{
				voucher: entity.Voucher{
					ID:         1,
					Code:       "ERRINVALIDEXPIRED",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: nil,
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			v, err := serviceVoucher.SaveVoucher(tc.args.voucher)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			vStr, wantVoucherStr := "nil", "nil"

			if v != nil {
				vStr = v.MarshalString()
			}

			if tc.wantVoucher != nil {
				wantVoucherStr = tc.wantVoucher.MarshalString()
			}

			if vStr != wantVoucherStr {
				t.Fatalf("want: %s\ngot: %s", wantVoucherStr, vStr)
			}

		})
	}
}

// go test -v -failfast -cover -run ^Test_GetVoucherByID$
func Test_GetVoucherByID(t *testing.T) {
	serviceVoucher := NewVoucherService(nil)

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
			err := serviceVoucher.GetVoucherByID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -failfast -cover -run ^Test_UpdateVoucherByID$
func Test_UpdateVoucherByID(t *testing.T) {
	v := entity.Voucher{}
	serviceVoucher := NewVoucherService(&v)

	type args struct {
		id      int64
		voucher entity.Voucher
	}

	tests := []struct {
		name        string
		args        args
		wantVoucher *entity.Voucher
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
				voucher: entity.Voucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: &entity.Voucher{
				ID:         1,
				Code:       "SUCCESS10",
				Porcentage: 10,
				CreatedAt: common.CreatedAt{
					Value: time.Time{},
				},
				ExpiresAt: common.ExpiresAt{
					Value: &time.Time{},
				},
			},
			wantError: false,
		},
		{
			name: "error_invalid_id",
			args: args{
				voucher: entity.Voucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_expires_at_value",
			args: args{
				id: 1,
				voucher: entity.Voucher{
					ID:         1,
					Code:       "ERRINVALIDEXPIRED",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: nil,
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_code",
			args: args{
				id: 1,
				voucher: entity.Voucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_upper",
			args: args{
				voucher: entity.Voucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: 1000,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_lower",
			args: args{
				id: 1,
				voucher: entity.Voucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: -1,
					CreatedAt: common.CreatedAt{
						Value: time.Time{},
					},
					ExpiresAt: common.ExpiresAt{
						Value: &time.Time{},
					},
				},
			},
			wantVoucher: nil,
			wantError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(*testing.T) {
			v, err := serviceVoucher.UpdateVoucherByID(tc.args.id, tc.args.voucher)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			vStr, wantVoucherStr := "nil", "nil"

			if v != nil {
				vStr = v.MarshalString()
			}

			if tc.wantVoucher != nil {
				wantVoucherStr = tc.wantVoucher.MarshalString()
			}

			if vStr != wantVoucherStr {
				t.Fatalf("want: %s\ngot: %s", wantVoucherStr, vStr)
			}

		})
	}
}
