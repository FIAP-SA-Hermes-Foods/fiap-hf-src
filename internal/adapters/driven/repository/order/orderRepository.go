package order

import (
	"context"
	"fiap-hf-src/internal/core/db"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"reflect"
)

var (
	queryGetOrders    = `SELECT * FROM orders ORDER BY created_at ASC`
	queryGetOrderByID = `SELECT * FROM orders WHERE id = $1`
	querySaveOrder    = `INSERT INTO orders (id, status, verification_code, created_at, client_id, voucher_id) VALUES (DEFAULT, $1, $2, now(), $3, $4) RETURNING id, created_at`
	queryUpdateOrder  = `UPDATE orders SET status = $1, client_id = $2, voucher_id = $3 WHERE id = $4 RETURNING id, created_at`
)

type OrderRepository interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
	GetOrderByID(id int64) (*entity.Order, error)
	GetOrders() ([]entity.Order, error)
	UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error)
}

type orderRepository struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewOrderRepository(ctx context.Context, db db.SQLDatabase) OrderRepository {
	return orderRepository{Ctx: ctx, Database: db}
}

func (o orderRepository) SaveOrder(order entity.Order) (*entity.Order, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveOrder); err != nil {
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &entity.Order{
		ClientID:  order.ClientID,
		VoucherID: order.VoucherID,
		Status: valueObject.Status{
			Value: order.Status.Value,
		},
		VerificationCode: valueObject.VerificationCode{
			Value: order.VerificationCode.Value,
		},
	}

	o.Database.QueryRow(order.Status.Value, order.VerificationCode.Value, order.ClientID, order.VoucherID)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outOrder, nil
}

func (o orderRepository) UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(queryUpdateOrder); err != nil {
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &entity.Order{
		ClientID:  order.ClientID,
		VoucherID: order.VoucherID,
		Status: valueObject.Status{
			Value: order.Status.Value,
		},
		VerificationCode: valueObject.VerificationCode{
			Value: order.VerificationCode.Value,
		},
	}

	o.Database.QueryRow(order.Status.Value, order.ClientID, order.VoucherID, id)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outOrder, nil
}

func (o orderRepository) GetOrderByID(id int64) (*entity.Order, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	var outOrder = new(entity.Order)

	if err := o.Database.Query(queryGetOrderByID, id); err != nil {
		return nil, err
	}

	for o.Database.GetNextRows() {
		err := o.Database.Scan(
			&outOrder.ID,
			&outOrder.Status.Value,
			&outOrder.VerificationCode.Value,
			&outOrder.CreatedAt.Value,
			&outOrder.ClientID,
			&outOrder.VoucherID,
		)
		if err != nil {
			return nil, err
		}
	}

	if reflect.ValueOf(outOrder).IsNil() || reflect.ValueOf(outOrder).IsZero() {
		return nil, nil
	}

	return outOrder, nil
}

func (o orderRepository) GetOrders() ([]entity.Order, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.Order)
		orderList = make([]entity.Order, 0)
	)

	if err := o.Database.Query(queryGetOrders); err != nil {
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem entity.Order

		err := o.Database.Scan(
			&order.ID,
			&order.Status.Value,
			&order.VerificationCode.Value,
			&order.CreatedAt.Value,
			&order.ClientID,
			&order.VoucherID,
		)

		if err != nil {
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	return orderList, nil
}
