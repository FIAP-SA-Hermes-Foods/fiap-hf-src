package rest

import (
	"fiap-hf-src/src/base/interfaces"
	"net/http"
)

var _ interfaces.VoucherHandler = (*handlerVoucher)(nil)

type handlerVoucher struct {
	controller interfaces.VoucherController
}

func NewHandlerVoucher(controller interfaces.VoucherController) *handlerVoucher {
	return &handlerVoucher{controller: controller}
}

func (h handlerVoucher) Handler(rw http.ResponseWriter, req *http.Request) {

	var routeVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/voucher/{id}": h.controller.GetVoucherByID,
		"post hermes_foods/voucher":     h.controller.SaveVoucher,
		"put hermes_foods/voucher/{id}": h.controller.UpdateVoucherByID,
	}

	handler, err := router(req.Method, req.URL.Path, routeVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}
