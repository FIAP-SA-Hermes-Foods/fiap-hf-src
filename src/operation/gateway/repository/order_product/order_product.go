package orderproduct

import (
	"context"
	"encoding/json"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
	l "fiap-hf-src/src/external/logger"
	ps "fiap-hf-src/src/operation/presenter/strings"
)

var (
	queryGetOrderProductByOrderID = `SELECT * from orders_products where orders_id = $1`
	queryGetOrderProducts         = `SELECT * from orders_products`
	querySaveOrderProducts        = `INSERT INTO orders_products (id, orders_id, product_id, quantity, total_price, discount, created_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now()) RETURNING id, created_at`
)

type OrderProductRepository interface {
	GetAllOrderProduct() ([]entity.OrderProduct, error)
	GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error)
	SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error)
}

type orderProductRepository struct {
	Ctx      context.Context
	Database interfaces.SQLDatabase
}

func NewOrderProductRepository(ctx context.Context, db interfaces.SQLDatabase) OrderProductRepository {
	return orderProductRepository{Ctx: ctx, Database: db}
}

func (o orderProductRepository) GetAllOrderProduct() ([]entity.OrderProduct, error) {
	l.Infof("GetAllOrderProduct received input: ", " | ", nil)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetAllOrderProduct connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.OrderProduct)
		orderList = make([]entity.OrderProduct, 0)
	)

	if err := o.Database.Query(queryGetOrderProducts); err != nil {
		l.Errorf("GetAllOrderProduct error to connect database: ", " | ", err)
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
			l.Errorf("GetAllOrderProduct error to scan database: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetAllOrderProductByOrderID error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetAllOrderProductByOrderID output: ", " | ", string(olStr))
	return orderList, nil
}

func (o orderProductRepository) GetAllOrderProductByOrderID(id int64) ([]entity.OrderProduct, error) {
	l.Infof("GetAllOrderProductByOrderID received input: ", " | ", id)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetAllOrderProductByOrderID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(entity.OrderProduct)
		orderList = make([]entity.OrderProduct, 0)
	)

	if err := o.Database.Query(queryGetOrderProductByOrderID, id); err != nil {
		l.Errorf("GetAllOrderProductByOrderID error to connect database: ", " | ", err)
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
			l.Errorf("GetAllOrderProductByOrderID error to scan database: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetAllOrderProductByOrderID error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetAllOrderProductByOrderID output: ", " | ", string(olStr))
	return orderList, nil
}

func (o orderProductRepository) SaveOrderProduct(order entity.OrderProduct) (*entity.OrderProduct, error) {
	l.Infof("SaveOrderProduct received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("SaveOrderProduct connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveOrderProducts); err != nil {
		l.Errorf("SaveOrderProduct error to prepare statement: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &entity.OrderProduct{
		Quantity:   order.ID,
		TotalPrice: order.TotalPrice,
		Discount:   order.Discount,
		OrderID:    order.OrderID,
		ProductID:  order.ProductID,
	}

	o.Database.QueryRow(order.OrderID, order.ProductID, order.Quantity, order.TotalPrice, order.Discount)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt.Value); err != nil {
		l.Errorf("SaveOrderProduct error to scan statement: ", " | ", err)
		return nil, err
	}

	l.Infof("SaveOrderProduct output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}
