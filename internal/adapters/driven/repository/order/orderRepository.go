package order

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	querySaveOrder = `INSERT INTO orders (id, status, verification_code, created_at, client_id, voucher_id) VALUES (DEFAULT, $1, $2, now(), $3, $4) RETURNING id, created_at`
)

type OrderRepository interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
}

type orderRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewOrderRepository(ctx context.Context, db psqldb.PostgresDB) OrderRepository {
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
