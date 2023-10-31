package application

import "fiap-hf-src/internal/core/domain/entity"

// Client implementation Call

func (app hermesFoodsApp) GetClientByCPFRepository(cpf string) (*entity.Client, error) {
	return app.clientRepo.GetClientByCPF(cpf)
}

func (app hermesFoodsApp) GetClientByIDRepository(id int64) (*entity.Client, error) {
	return app.clientRepo.GetClientByID(id)
}

func (app hermesFoodsApp) SaveClientRepository(client entity.Client) (*entity.Client, error) {
	return app.clientRepo.SaveClient(client)
}

// Order implementation Call

func (app hermesFoodsApp) GetOrdersRepository() ([]entity.Order, error) {
	return app.orderRepo.GetOrders()
}

func (app hermesFoodsApp) GetOrderByIDRepository(id int64) (*entity.Order, error) {
	return app.orderRepo.GetOrderByID(id)
}

func (app hermesFoodsApp) SaveOrderRepository(order entity.Order) (*entity.Order, error) {
	return app.orderRepo.SaveOrder(order)
}

func (app hermesFoodsApp) UpdateOrderByIDRepository(id int64, order entity.Order) (*entity.Order, error) {
	return app.orderRepo.UpdateOrderByID(id, order)
}

// OrderProduct implementation Call

func (app hermesFoodsApp) GetAllOrderProduct() ([]entity.OrderProduct, error) {
	return app.orderProductRepo.GetAllOrderProduct()
}

func (app hermesFoodsApp) GetAllOrderProductByIdRepository(id int64) ([]entity.OrderProduct, error) {
	return app.orderProductRepo.GetAllOrderProductByOrderID(id)
}

func (app hermesFoodsApp) SaveOrderProductRepository(orderProduct entity.OrderProduct) (*entity.OrderProduct, error) {
	return app.orderProductRepo.SaveOrderProduct(orderProduct)
}

// Product implementation Call

func (app hermesFoodsApp) GetProductByIDRepository(id int64) (*entity.Product, error) {
	return app.productRepo.GetProductByID(id)
}

func (app hermesFoodsApp) GetProductByCategoryRepository(category string) ([]entity.Product, error) {
	return app.productRepo.GetProductByCategory(category)
}

func (app hermesFoodsApp) SaveProductRepository(product entity.Product) (*entity.Product, error) {
	return app.productRepo.SaveProduct(product)
}

func (app hermesFoodsApp) UpdateProductByIDRepository(id int64, product entity.Product) (*entity.Product, error) {
	return app.productRepo.UpdateProductByID(id, product)
}

func (app hermesFoodsApp) DeleteProductByIDRepository(id int64) error {
	return app.productRepo.DeleteProductByID(id)
}

// Voucher implementation Call

func (app hermesFoodsApp) GetVoucherByIDRepository(id int64) (*entity.Voucher, error) {
	return app.voucherRepo.GetVoucherByID(id)
}

func (app hermesFoodsApp) SaveVoucherRepository(voucher entity.Voucher) (*entity.Voucher, error) {
	return app.voucherRepo.SaveVoucher(voucher)
}

func (app hermesFoodsApp) UpdateVoucherByIDRepository(id int64, voucher entity.Voucher) (*entity.Voucher, error) {
	return app.voucherRepo.UpdateVoucherByID(id, voucher)
}
