package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type (
	PostgresDB interface {
		Connect() error
		Close() error
		PrepareStmt(query string) error

		/*
		   	ExecContext: This function will query a prepared statement, and return its result

		   IMPORTANT!:
		     - This method only works after running the method: *PrepareStmt*
		*/
		ExecContext(ctx context.Context, fields ...interface{}) (sql.Result, error)

		/*
		   	Query: This function will query a prepared statement and return its rows

		   IMPORTANT!:
		     - This method only works after running the method: *PrepareStmt*
		*/
		Query(args ...interface{}) (*sql.Rows, error)

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
		Scan(args ...interface{}) error
	}
)

type (
	postgresDB struct {
		Ctx        context.Context
		Host       string
		Port       string
		Schema     string
		User       string
		Password   string
		postgresDB *sql.DB
		SqlStmt    *sql.Stmt
		Row        *sql.Row
	}
)

func NewPostgresDB(ctx context.Context, host, port, schema, user, password string) PostgresDB {
	return &postgresDB{
		Ctx:      ctx,
		Host:     host,
		Port:     port,
		Schema:   schema,
		User:     user,
		Password: password,
	}
}

// DB Methods

func (p *postgresDB) Connect() error {
	if p.postgresDB != nil {
		return nil
	}

	db, err := sql.Open("postgres", p.dbURL())

	if err != nil {
		return err
	}

	if err := p.PingCtx(p.Ctx); err != nil {
		return nil
	}

	p.postgresDB = db

	return nil
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
	if p.SqlStmt != nil {
		return nil
	}

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
	url := fmt.Sprintf("%s:%s", p.Host, p.Port)
	return fmt.Sprintf("postgresql://%s:%s@%s/%s/sslmode=disable", p.User, p.Password, url, p.Schema)
}

// Stmt Methods

/*
	Query: This function will query a prepared statement and return its rows

IMPORTANT!:
  - This method only works after running the method: *PrepareStmt*
*/
func (p *postgresDB) Query(args ...interface{}) (*sql.Rows, error) {
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
	if p.Row != nil {
		return
	}

	if p.SqlStmt != nil {
		p.Row = p.SqlStmt.QueryRow(args...)
	}
}

/*
	ExecContext: This function will query a prepared statement, and return its result

IMPORTANT!:
  - This method only works after running the method: *PrepareStmt*
*/
func (p *postgresDB) ExecContext(ctx context.Context, fields ...interface{}) (sql.Result, error) {
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
func (s *postgresDB) Scan(args ...interface{}) error {
	if s.Row != nil {
		return s.Row.Scan(args...)
	}

	return errors.New("row is null, is not possible to scan")
}
