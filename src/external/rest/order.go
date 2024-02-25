package rest

import (
	"fiap-hf-src/src/base/interfaces"
	"net/http"
)

var _ interfaces.OrderHandler = (*handlerOrder)(nil)

type handlerOrder struct {
	controller interfaces.OrderController
}

func NewHandlerOrder(controller interfaces.OrderController) *handlerOrder {
	return &handlerOrder{controller: controller}
}

func (h *handlerOrder) Handler(rw http.ResponseWriter, req *http.Request) {

	apiHToken := req.Header.Get("Auth-token")

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	var routesOrders = map[string]http.HandlerFunc{
		"get hermes_foods/order":        h.controller.GetOrders,
		"get hermes_foods/order/{id}":   h.controller.GetOrderByID,
		"post hermes_foods/order":       h.controller.SaveOrder,
		"patch hermes_foods/order/{id}": h.controller.UpdateOrderByID,
	}

	handler, err := router(req.Method, req.URL.Path, routesOrders)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}
