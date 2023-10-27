package migration

import (
	"context"
	"database/sql"
	"testing"
)

// go test -v -failfast -count=1 -cover -run ^Test_DBMigration$
func Test_DBMigration(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name    string
		ctx     context.Context
		mockdb  *mockDb
		wantErr bool
	}{
		{
			name: "success",
			ctx:  ctx,
			mockdb: &mockDb{
				WantResult: nil,
				WantRows:   &sql.Rows{},
				WantErr:    nil,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			migration := NewMigration(tc.mockdb)

			if err := migration.Migrate(); err != nil {
				if !tc.wantErr {
					t.Errorf("was suppose to not have an error here and got:\n%v", err)
				}
			}
		})
	}
}
