package client

import (
	"context"
	"database/sql"
	"strings"
)

// Mock DB

type mockDb struct {
	WantResult sql.Result
	WantRows   *sql.Rows
	WantErr    error
}

func (m mockDb) Connect() error {
	if m.WantErr != nil && strings.EqualFold("errConnect", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m mockDb) Close() error {
	if m.WantErr != nil && strings.EqualFold("errClose", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m mockDb) PrepareStmt(query string) error {
	if m.WantErr != nil && strings.EqualFold("errPrepareStmt", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}

func (m mockDb) ExecContext(ctx context.Context, fields ...interface{}) (sql.Result, error) {
	if m.WantErr != nil && strings.EqualFold("errExecContext", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResult, nil
}

func (m mockDb) Query(args ...interface{}) (*sql.Rows, error) {
	if m.WantErr != nil && strings.EqualFold("errQuery", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantRows, nil
}

func (m mockDb) QueryRow(args ...interface{}) {
}

func (m mockDb) CloseStmt() error {
	if m.WantErr != nil && strings.EqualFold("errCloseStmt", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil

}

func (m mockDb) Scan(args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}
