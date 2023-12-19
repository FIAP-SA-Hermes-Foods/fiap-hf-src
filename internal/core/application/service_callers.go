package application

import "fiap-hf-src/internal/core/entity"

// Client implementation Call

func (app hermesFoodsApp) GetClientByCPFService(cpf string) error {
	return app.clientService.GetClientByCPF(cpf)
}

func (app hermesFoodsApp) GetClientByIDService(id int64) error {
	return app.clientService.GetClientByID(id)
}

func (app hermesFoodsApp) SaveClientService(client entity.Client) (*entity.Client, error) {
	return app.clientService.SaveClient(client)
}

// Order implementation Call

func (app hermesFoodsApp) GetOrderByIDService(id int64) error {
	return app.orderService.GetOrderByID(id)
}

func (app hermesFoodsApp) SaveOrderService(order entity.Order) (*entity.Order, error) {
	return app.orderService.SaveOrder(order)
}

func (app hermesFoodsApp) UpdateOrderByIDService(id int64, order entity.Order) (*entity.Order, error) {
	return app.orderService.UpdateOrderByID(id, order)
}

// OrderProduct implementation Call

func (app hermesFoodsApp) GetAllOrderProductByIdService(id int64) error {
	return app.orderProductService.GetOrderProductByOrderID(id)
}

func (app hermesFoodsApp) SaveOrderProductService(orderProduct entity.OrderProduct) (*entity.OrderProduct, error) {
	return app.orderProductService.SaveOrderProduct(orderProduct)
}

// Product implementation Call

func (app hermesFoodsApp) SaveProductService(product entity.Product) (*entity.Product, error) {
	return app.productService.SaveProduct(product)
}

func (app hermesFoodsApp) GetProductByIDService(id int64) error {
	return app.productService.GetProductByID(id)
}

func (app hermesFoodsApp) GetProductByCategoryService(category string) error {
	return app.productService.GetProductByCategory(category)
}

func (app hermesFoodsApp) UpdateProductByIDService(id int64, product entity.Product) (*entity.Product, error) {
	return app.productService.UpdateProductByID(id, product)
}

func (app hermesFoodsApp) DeleteProductByIDService(id int64) error {
	return app.productService.DeleteProductByID(id)
}

// Voucher implementation Call

func (app hermesFoodsApp) GetVoucherByIDService(id int64) error {
	return app.voucherService.GetVoucherByID(id)
}

func (app hermesFoodsApp) SaveVoucherService(voucher entity.Voucher) (*entity.Voucher, error) {
	return app.voucherService.SaveVoucher(voucher)
}

func (app hermesFoodsApp) UpdateVoucherByIDService(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
	return app.voucherService.UpdateVoucherByID(id, voucher)
}
