package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fiap-hf-src/src/base/interfaces"
	"fmt"
	"time"
)

var _ interfaces.SQLDatabase = (*postgresDB)(nil)

type postgresDB struct {
	Ctx        context.Context
	Region     string // us-east-1
	Host       string // my-postgres-db-instance.123456789012.us-east-1.rds.amazonaws.com
	Port       string
	Schema     string
	User       string
	Password   string
	Timeout    time.Duration
	postgresDB *sql.DB
	SqlStmt    *sql.Stmt
	Row        *sql.Row
	Rows       *sql.Rows
}

func NewPostgresDB(ctx context.Context, region, host, port, schema, user, password string, timeout time.Duration) *postgresDB {
	return &postgresDB{
		Ctx:      ctx,
		Region:   region,
		Host:     host,
		Port:     port,
		Schema:   schema,
		User:     user,
		Password: password,
		Timeout:  timeout,
	}
}

// DB Methods
func (p *postgresDB) Connect() error {

	db, err := sql.Open("postgres", p.dbURL())

	if err != nil {
		return err
	}

	if err := db.PingContext(p.Ctx); err != nil {
		return err
	}

	p.postgresDB = db

	return nil
}

func (p *postgresDB) ExecContext(ctx context.Context, query string, fields ...interface{}) (sql.Result, error) {
	if p.postgresDB != nil {
		return p.postgresDB.ExecContext(ctx, query, fields...)
	}

	return nil, errors.New("connection is null, is not possible to ExecContext")
}

func (p *postgresDB) Query(query string, args ...interface{}) error {
	if p.postgresDB == nil {
		return errors.New("connection is null, is not possible to Query")
	}

	rows, err := p.postgresDB.Query(query, args...)
	if err != nil {
		return err
	}

	p.Rows = rows

	return nil
}

func (s *postgresDB) Scan(args ...interface{}) error {
	if s.Rows != nil {
		return s.Rows.Scan(args...)
	}

	return errors.New("row is null, is not possible to scan")
}

func (s *postgresDB) GetNextRows() bool {
	return s.Rows.Next()
}

func (p *postgresDB) PingCtx(ctx context.Context) error {
	if p.postgresDB != nil {
		return p.postgresDB.PingContext(ctx)
	}
	return errors.New("connection is null, is not possible to ping")
}

func (p *postgresDB) Close() error {
	if p.postgresDB != nil {
		return p.postgresDB.Close()
	}

	return errors.New("connection is null, is not possible to close it")
}

func (p *postgresDB) PrepareStmt(query string) error {

	if p.postgresDB == nil {
		return errors.New("connection is null, is not possible to create a stmt")
	}

	s, err := p.postgresDB.PrepareContext(p.Ctx, query)
	if err != nil {
		return err
	}

	p.SqlStmt = s

	return nil
}

func (p postgresDB) dbURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", p.User, p.Password, p.Host, p.Port, p.Schema)
}

// Stmt Methods

/*
	Query: This function will query a prepared statement and return its rows

IMPORTANT!:
  - This method only works after running the method: *PrepareStmt*
*/
func (p *postgresDB) QueryStmt(args ...interface{}) (*sql.Rows, error) {
	if p.SqlStmt != nil {
		return p.SqlStmt.Query(args...)
	}
	return nil, errors.New("stmt is null, is not possible to query")
}

/*
	QueryRow: This function will query a prepared statement

IMPORTANT!:
  - This method only works after running the method: *PrepareStmt*
*/
func (p *postgresDB) QueryRow(args ...interface{}) {
	if p.SqlStmt != nil {
		p.Row = p.SqlStmt.QueryRow(args...)
	}
}

/*
	ExecContext: This function will query a prepared statement, and return its result

IMPORTANT!:
  - This method only works after running the method: *PrepareStmt*
*/
func (p *postgresDB) ExecContextStmt(ctx context.Context, fields ...interface{}) (sql.Result, error) {
	if p.SqlStmt != nil {
		return p.SqlStmt.ExecContext(ctx, fields...)
	}

	return nil, errors.New("stmt is null")
}

func (p *postgresDB) CloseStmt() error {
	if p.SqlStmt != nil {
		return p.SqlStmt.Close()
	}
	return errors.New("connection is null, is not possible to close stmt")
}

// Row Methods

/*
	Scan: This method scans all args in input and provide values to them through a executed sql script

IMPORTANT!:
  - This method only works after running the method: *QueryRow*
*/
func (s *postgresDB) ScanStmt(args ...interface{}) error {
	if s.Row != nil {
		return s.Row.Scan(args...)
	}

	return errors.New("row is null, is not possible to scan")
}
