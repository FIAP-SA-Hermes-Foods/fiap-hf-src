package orderproduct

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
)

var (
	queryGetOrderProductByOrderID = `SELECT * from orders_products where orders_id = $1`
	queryGetOrderProducts         = `SELECT * from orders_products`
)

type OrderProductRepository interface {
	GetAllOrderProduct() ([]entity.OrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error)
}

type orderProductRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewOrderProductRepository(ctx context.Context, db psqldb.PostgresDB) OrderProductRepository {
	return orderProductRepository{Ctx: ctx, Database: db}
}

func (o orderProductRepository) GetAllOrderProduct() ([]entity.OrderProduct, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.OrderProduct)
		orderList = make([]entity.OrderProduct, 0)
	)

	if err := o.Database.Query(queryGetOrderProducts); err != nil {
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem entity.OrderProduct

		err := o.Database.Scan(
			&order.ID,
			&order.Quantity,
			&order.TotalPrice,
			&order.Discount,
			&order.OrderID,
			&order.ProductID,
			&order.CreatedAt.Value,
		)

		if err != nil {
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	return orderList, nil
}

func (o orderProductRepository) GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.OrderProduct)
		orderList = make([]entity.OrderProduct, 0)
	)

	if err := o.Database.Query(queryGetOrderProductByOrderID, id); err != nil {
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem entity.OrderProduct

		err := o.Database.Scan(
			&order.ID,
			&order.Quantity,
			&order.TotalPrice,
			&order.Discount,
			&order.OrderID,
			&order.ProductID,
			&order.CreatedAt.Value,
		)

		if err != nil {
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	return orderList, nil
}
