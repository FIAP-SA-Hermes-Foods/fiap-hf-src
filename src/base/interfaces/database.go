package interfaces

import (
	"context"
	"database/sql"
)

type SQLDatabase interface {
	Connect() error
	Close() error
	PrepareStmt(query string) error
	ExecContext(ctx context.Context, query string, fields ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) error
	GetNextRows() bool
	Scan(args ...interface{}) error
	/*
	   	ExecContext: This function will query a prepared statement, and return its result

	   IMPORTANT!:
	     - This method only works after running the method: *PrepareStmt*
	*/
	ExecContextStmt(ctx context.Context, fields ...interface{}) (sql.Result, error)

	/*
	   	Query: This function will query a prepared statement and return its rows

	   IMPORTANT!:
	     - This method only works after running the method: *PrepareStmt*
	*/
	QueryStmt(args ...interface{}) (*sql.Rows, error)

	/*
	   	QueryRow: This function will query a prepared statement

	   IMPORTANT!:
	     - This method only works after running the method: *PrepareStmt*
	*/
	QueryRow(args ...interface{})
	CloseStmt() error

	/*
	   	Scan: This method scans all args in input and provide values to them through a executed sql script

	   IMPORTANT!:
	     - This method only works after running the method: *QueryRow*
	*/
	ScanStmt(args ...interface{}) error
}
