package interfaces

import "net/http"

type ClientHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type OrderHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type ProductHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type VoucherHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}
