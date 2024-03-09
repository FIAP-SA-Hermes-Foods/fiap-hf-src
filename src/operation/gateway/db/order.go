package db

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
)

var _ interfaces.OrderGateway = (*orderGateway)(nil)

type orderGateway struct {
	db interfaces.OrderDB
}

func NewOrderGateway(db interfaces.OrderDB) *orderGateway {
	return &orderGateway{db: db}
}

func (o *orderGateway) SaveOrder(reqOrder dto.RequestOrder) (*dto.OutputOrder, error) {
	order := reqOrder.Order()

	oDb, err := o.db.SaveOrder(order)

	if err != nil {
		return nil, err
	}

	if oDb == nil {
		return nil, nil
	}

	products := make([]entity.ProductItem, 0)

	for i := range oDb.Items {
		p := entity.ProductItem{
			ID: oDb.Items[i].ProductID,
		}
		products = append(products, p)
	}

	out := &dto.OutputOrder{
		ID: oDb.ID,
		Client: dto.OutputClient{
			ID: order.ClientID,
		},
		VoucherID:        oDb.VoucherID,
		VerificationCode: oDb.VerificationCode.Value,
		CreatedAt:        oDb.CreatedAt.Format(),
		Status:           oDb.Status.Value,
		Products:         products,
	}

	return out, nil
}

func (o *orderGateway) UpdateOrderByID(id int64, reqOrder dto.RequestOrder) (*dto.OutputOrder, error) {
	order := reqOrder.Order()

	oDb, err := o.db.UpdateOrderByID(id, order)

	if err != nil {
		return nil, nil
	}

	if oDb == nil {
		return nil, nil
	}

	products := make([]entity.ProductItem, 0)

	for i := range oDb.Items {
		p := entity.ProductItem{
			ID: oDb.Items[i].ProductID,
		}
		products = append(products, p)
	}

	out := &dto.OutputOrder{
		ID: oDb.ID,
		Client: dto.OutputClient{
			ID: oDb.ClientID,
		},
		VoucherID:        oDb.VoucherID,
		VerificationCode: oDb.VerificationCode.Value,
		CreatedAt:        oDb.CreatedAt.Format(),
		Status:           oDb.Status.Value,
		Products:         products,
	}

	return out, nil
}

func (o *orderGateway) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	order, err := o.db.GetOrderByID(id)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	products := make([]entity.ProductItem, 0)

	for i := range order.Items {
		p := entity.ProductItem{
			ID: order.Items[i].ProductID,
		}
		products = append(products, p)
	}

	out := &dto.OutputOrder{
		ID: order.ID,
		Client: dto.OutputClient{
			ID: order.ClientID,
		},
		VoucherID:        order.VoucherID,
		VerificationCode: order.VerificationCode.Value,
		CreatedAt:        order.CreatedAt.Format(),
		Status:           order.Status.Value,
		Products:         products,
	}
	return out, nil
}

func (o *orderGateway) GetOrders() ([]dto.OutputOrder, error) {

	oDb, err := o.db.GetOrders()

	if err != nil {
		return nil, err
	}
	if oDb == nil {
		return nil, nil
	}

	var out = make([]dto.OutputOrder, 0)

	for i := range oDb {
		items := make([]entity.ProductItem, 0)

		for p := range oDb[i].Items {
			pi := entity.ProductItem{
				ID: oDb[i].Items[p].ProductID,
			}

			items = append(items, pi)
		}

		o := dto.OutputOrder{
			ID: oDb[i].ID,
			Client: dto.OutputClient{
				ID: oDb[i].ClientID,
			},
			Products:         items,
			VoucherID:        oDb[i].VoucherID,
			Status:           oDb[i].Status.Value,
			VerificationCode: oDb[i].VerificationCode.Value,
			CreatedAt:        oDb[i].CreatedAt.Format(),
		}

		out = append(out, o)
	}

	return out, nil
}
