package interfaces

import "net/http"

type ClientController interface {
	GetClientByCPF(rw http.ResponseWriter, req *http.Request)
	SaveClient(rw http.ResponseWriter, req *http.Request)
}

type OrderController interface {
	GetOrders(rw http.ResponseWriter, req *http.Request)
	GetOrderByID(rw http.ResponseWriter, req *http.Request)
	SaveOrder(rw http.ResponseWriter, req *http.Request)
	UpdateOrderByID(rw http.ResponseWriter, req *http.Request)
}

type ProductController interface {
	GetProductByCategory(rw http.ResponseWriter, req *http.Request)
	SaveProduct(rw http.ResponseWriter, req *http.Request)
	UpdateProductByID(rw http.ResponseWriter, req *http.Request)
	DeleteProductByID(rw http.ResponseWriter, req *http.Request)
}

type VoucherController interface {
	GetVoucherByID(rw http.ResponseWriter, req *http.Request)
	SaveVoucher(rw http.ResponseWriter, req *http.Request)
	UpdateVoucherByID(rw http.ResponseWriter, req *http.Request)
}
