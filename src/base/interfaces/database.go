package interfaces

import (
	"context"
	"database/sql"
	"fiap-hf-src/src/core/entity"
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

type ClientDB interface {
	GetClientByID(id int64) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type OrderDB interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
	GetOrders() ([]entity.Order, error)
	GetOrderByID(id int64) (*entity.Order, error)
}

type OrderProductDB interface {
	GetAllOrderProduct() ([]entity.OrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error)
	SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error)
}

type ProductDB interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByCategory(category string) ([]entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	DeleteProductByID(id int64) error
}

type VoucherDB interface {
	GetVoucherByID(id int64) (*entity.Voucher, error)
	SaveVoucher(voucher entity.Voucher) (*entity.Voucher, error)
	UpdateVoucherByID(id int64, voucher entity.Voucher) (*entity.Voucher, error)
}
