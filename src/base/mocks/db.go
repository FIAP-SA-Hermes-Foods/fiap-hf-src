package mocks

import (
	"context"
	"database/sql"
	"strings"
)

// Mock DB

type MockDb struct {
	WantResult   sql.Result
	WantRows     *sql.Rows
	WantErr      error
	WantNextRows bool
}

func (m MockDb) Connect() error {
	if m.WantErr != nil && strings.EqualFold("errConnect", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDb) Close() error {
	if m.WantErr != nil && strings.EqualFold("errClose", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDb) PrepareStmt(query string) error {
	if m.WantErr != nil && strings.EqualFold("errPrepareStmt", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDb) GetNextRows() bool {
	return m.WantNextRows
}

func (m MockDb) ExecContext(ctx context.Context, query string, fields ...interface{}) (sql.Result, error) {
	if m.WantErr != nil && strings.EqualFold("errExecContext", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResult, nil
}

func (m MockDb) ExecContextStmt(ctx context.Context, fields ...interface{}) (sql.Result, error) {
	if m.WantErr != nil && strings.EqualFold("errExecContextStmt", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResult, nil
}

func (m MockDb) Query(query string, args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errQuery", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDb) QueryStmt(args ...interface{}) (*sql.Rows, error) {
	if m.WantErr != nil && strings.EqualFold("errQueryStmt", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantRows, nil
}

func (m MockDb) QueryRow(args ...interface{}) {
}

func (m MockDb) CloseStmt() error {
	if m.WantErr != nil && strings.EqualFold("errCloseStmt", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil

}

func (m MockDb) Scan(args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}

func (m MockDb) ScanStmt(args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errScanStmt", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}
