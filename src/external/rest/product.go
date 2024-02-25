package rest

import (
	"fiap-hf-src/src/base/interfaces"
	"net/http"
)

var _ interfaces.ProductHandler = (*handlerProduct)(nil)

type handlerProduct struct {
	controller interfaces.ProductController
}

func NewHandlerProduct(controller interfaces.ProductController) *handlerProduct {
	return &handlerProduct{controller: controller}
}

func (h handlerProduct) Handler(rw http.ResponseWriter, req *http.Request) {
	apiHToken := req.Header.Get("Auth-token")

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	var routeProducts = map[string]http.HandlerFunc{
		"get hermes_foods/product":         h.controller.GetProductByCategory,
		"post hermes_foods/product":        h.controller.SaveProduct,
		"put hermes_foods/product/{id}":    h.controller.UpdateProductByID,
		"delete hermes_foods/product/{id}": h.controller.DeleteProductByID,
	}

	handler, err := router(req.Method, req.URL.Path, routeProducts)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}
