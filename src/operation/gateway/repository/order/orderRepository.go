package order

import (
	"context"
	"encoding/json"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
	l "fiap-hf-src/src/external/logger"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
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
	Database interfaces.SQLDatabase
}

func NewOrderRepository(ctx context.Context, db interfaces.SQLDatabase) OrderRepository {
	return orderRepository{Ctx: ctx, Database: db}
}

func (o orderRepository) SaveOrder(order entity.Order) (*entity.Order, error) {
	l.Infof("SaveOrder received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("SaveOrder connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveOrder); err != nil {
		l.Errorf("SaveOrder prepare error: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &entity.Order{
		ClientID:  order.ClientID,
		VoucherID: order.VoucherID,
		Status: com.Status{
			Value: order.Status.Value,
		},
		VerificationCode: com.VerificationCode{
			Value: order.VerificationCode.Value,
		},
	}

	o.Database.QueryRow(order.Status.Value, order.VerificationCode.Value, order.ClientID, order.VoucherID)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt.Value); err != nil {
		l.Errorf("SaveOrder scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("SaveOrder output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderRepository) UpdateOrderByID(id int64, order entity.Order) (*entity.Order, error) {
	l.Infof("UpdateOrderByID received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("UpdateOrderByID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(queryUpdateOrder); err != nil {
		l.Errorf("UpdateOrderByID prepare error: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &entity.Order{
		ClientID:  order.ClientID,
		VoucherID: order.VoucherID,
		Status: com.Status{
			Value: order.Status.Value,
		},
		VerificationCode: com.VerificationCode{
			Value: order.VerificationCode.Value,
		},
	}

	o.Database.QueryRow(order.Status.Value, order.ClientID, order.VoucherID, id)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt.Value); err != nil {
		l.Errorf("UpdateOrderByID scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("UpdateOrderByID output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderRepository) GetOrderByID(id int64) (*entity.Order, error) {
	l.Infof("GetOrderByID received input: ", " | ", id)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetOrderByID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var outOrder = new(entity.Order)

	if err := o.Database.Query(queryGetOrderByID, id); err != nil {
		l.Errorf("GetOrderByID query error: ", " | ", err)
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
			l.Errorf("GetOrderByID scan error: ", " | ", err)
			return nil, err
		}
	}

	if reflect.ValueOf(outOrder).IsNil() || reflect.ValueOf(outOrder).IsZero() {
		return nil, nil
	}

	l.Infof("GetOrderByID output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderRepository) GetOrders() ([]entity.Order, error) {
	l.Infof("GetOrders received input: ", " | ", nil)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetOrders connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.Order)
		orderList = make([]entity.Order, 0)
	)

	if err := o.Database.Query(queryGetOrders); err != nil {
		l.Errorf("GetOrders query error: ", " | ", err)
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
			l.Errorf("GetOrders scan error: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetOrders error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetOrders output: ", " | ", string(olStr))
	return orderList, nil
}
