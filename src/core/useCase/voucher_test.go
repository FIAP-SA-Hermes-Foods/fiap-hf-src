package useCase

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/mocks"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"testing"
)

// go test -v -failfast -cover -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {

	type args struct {
		voucher dto.RequestVoucher
	}

	tests := []struct {
		name        string
		args        args
		wantVoucher *dto.OutputVoucher
		mockGateway *mocks.VoucherGateway
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
					CreatedAt:  "10-01-2024 15:04:05",
					ExpiresAt:  "15-01-2025 15:04:05",
				},
			},
			wantVoucher: &dto.OutputVoucher{
				ID:         1,
				Code:       "SUCCESS10",
				Porcentage: 10,
			},
			mockGateway: &mocks.VoucherGateway{
				WantOut: &dto.OutputVoucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
				},
				WantErr: nil,
			},
			wantError: false,
		},
		{
			name: "error_invalid_code",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_upper",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: 1000,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_lower",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: -1,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_expires_at_value",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         1,
					Code:       "ERRINVALIDEXPIRED",
					Porcentage: 10,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
	}

	for _, tc := range tests {

		serviceVoucher := NewVoucherUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			v, err := serviceVoucher.SaveVoucher(tc.args.voucher)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			vStr, wantVoucherStr := "nil", "nil"

			if v != nil {
				vStr = ps.MarshalString(v)
			}

			if tc.wantVoucher != nil {
				wantVoucherStr = ps.MarshalString(tc.wantVoucher)
			}

			if vStr != wantVoucherStr {
				t.Fatalf("want: %s\ngot: %s", wantVoucherStr, vStr)
			}

		})
	}
}

// go test -v -failfast -cover -run ^Test_GetVoucherByID$
func Test_GetVoucherByID(t *testing.T) {

	type args struct {
		id int64
	}

	tests := []struct {
		name        string
		args        args
		wantError   bool
		mockGateway *mocks.VoucherGateway
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockGateway: &mocks.VoucherGateway{},
			wantError:   false,
		},
		{
			name: "error_invalid_id",
			args: args{
				id: 0,
			},
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceVoucher := NewVoucherUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			_, err := serviceVoucher.GetVoucherByID(tc.args.id)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}
		})
	}
}

// go test -v -failfast -cover -run ^Test_UpdateVoucherByID$
func Test_UpdateVoucherByID(t *testing.T) {

	type args struct {
		id      int64
		voucher dto.RequestVoucher
	}

	tests := []struct {
		name        string
		args        args
		wantVoucher *dto.OutputVoucher
		mockGateway *mocks.VoucherGateway
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
				voucher: dto.RequestVoucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
				},
			},
			wantVoucher: &dto.OutputVoucher{
				ID:         1,
				Code:       "SUCCESS10",
				Porcentage: 10,
			},
			mockGateway: &mocks.VoucherGateway{
				WantOut: &dto.OutputVoucher{
					ID:         1,
					Code:       "SUCCESS10",
					Porcentage: 10,
				},
				WantErr: nil,
			},
			wantError: false,
		},
		{
			name: "error_invalid_id",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_expires_at_value",
			args: args{
				id: 1,
				voucher: dto.RequestVoucher{
					ID:         1,
					Code:       "ERRINVALIDEXPIRED",
					Porcentage: 10,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_code",
			args: args{
				id: 1,
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "",
					Porcentage: 10,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_upper",
			args: args{
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: 1000,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
		{
			name: "error_invalid_porcentage_lower",
			args: args{
				id: 1,
				voucher: dto.RequestVoucher{
					ID:         0,
					Code:       "ERRPORCENTAGE10",
					Porcentage: -1,
				},
			},
			wantVoucher: nil,
			mockGateway: &mocks.VoucherGateway{},
			wantError:   true,
		},
	}

	for _, tc := range tests {
		serviceVoucher := NewVoucherUseCase(tc.mockGateway)
		t.Run(tc.name, func(*testing.T) {
			v, err := serviceVoucher.UpdateVoucherByID(tc.args.id, tc.args.voucher)

			if (!tc.wantError) && err != nil {
				t.Fatalf("error: %v", err)
			}

			vStr, wantVoucherStr := "nil", "nil"

			if v != nil {
				vStr = ps.MarshalString(v)
			}

			if tc.wantVoucher != nil {
				wantVoucherStr = ps.MarshalString(tc.wantVoucher)
			}

			if vStr != wantVoucherStr {
				t.Fatalf("want: %s\ngot: %s", wantVoucherStr, vStr)
			}
		})
	}
}
