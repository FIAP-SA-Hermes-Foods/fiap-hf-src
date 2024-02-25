package db

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.OrderProductGateway = (*orderProductGateway)(nil)

type orderProductGateway struct {
	db interfaces.OrderProductDB
}

func NewOrderProductGateway(db interfaces.OrderProductDB) *orderProductGateway {
	return &orderProductGateway{db: db}
}

func (o *orderProductGateway) GetAllOrderProduct() ([]dto.OutputOrderProduct, error) {
	out := make([]dto.OutputOrderProduct, 0)

	outDB, err := o.db.GetAllOrderProduct()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(outDB); i++ {
		o := dto.OutputOrderProduct{
			ID:         outDB[i].ID,
			Quantity:   outDB[i].Quantity,
			TotalPrice: outDB[i].TotalPrice,
			Discount:   outDB[i].Discount,
			OrderID:    outDB[i].OrderID,
			ProductID:  outDB[i].ProductID,
			CreatedAt:  outDB[i].CreatedAt.Format(),
		}
		out = append(out, o)
	}

	return out, nil
}

func (o *orderProductGateway) GetAllOrderProductByOrderID(id int64) ([]dto.OutputOrderProduct, error) {
	out := make([]dto.OutputOrderProduct, 0)

	outDB, err := o.db.GetAllOrderProductByOrderID(id)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(outDB); i++ {
		o := dto.OutputOrderProduct{
			ID:         outDB[i].ID,
			Quantity:   outDB[i].Quantity,
			TotalPrice: outDB[i].TotalPrice,
			Discount:   outDB[i].Discount,
			OrderID:    outDB[i].OrderID,
			ProductID:  outDB[i].ProductID,
			CreatedAt:  outDB[i].CreatedAt.Format(),
		}
		out = append(out, o)
	}

	return out, nil

}

func (o *orderProductGateway) SaveOrderProduct(reqOrder dto.RequestOrderProduct) (*dto.OutputOrderProduct, error) {

	order := reqOrder.OrderProduct()

	orderProduct, err := o.db.SaveOrderProduct(order)

	if err != nil {
		return nil, err
	}

	if orderProduct == nil {
		return nil, nil
	}

	out := &dto.OutputOrderProduct{
		ID:         orderProduct.ID,
		Quantity:   orderProduct.Quantity,
		TotalPrice: orderProduct.TotalPrice,
		Discount:   orderProduct.Discount,
		OrderID:    orderProduct.OrderID,
		ProductID:  orderProduct.ProductID,
		CreatedAt:  orderProduct.CreatedAt.Format(),
	}

	return out, nil

}
