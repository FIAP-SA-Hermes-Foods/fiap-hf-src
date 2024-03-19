package rest

import (
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fmt"
	"net/http"
)

var _ interfaces.ProductHandler = (*handlerProduct)(nil)

type handlerProduct struct {
	controller interfaces.ProductController
	userAuth   interfaces.UserAuth
}

func NewHandlerProduct(controller interfaces.ProductController, userAuth interfaces.UserAuth) *handlerProduct {
	return &handlerProduct{controller: controller, userAuth: userAuth}
}

func (h handlerProduct) Handler(rw http.ResponseWriter, req *http.Request) {

	userAuthStr := req.Header.Get("user-auth")

	if len(userAuthStr) > 0 {

		var uAuth dto.UserInput

		if err := json.Unmarshal([]byte(userAuthStr), &uAuth); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "%v"} `, err)
			return
		}

		if uAuth.WantRegister {
			out, err := h.userAuth.Auth(uAuth)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(rw, `{"error": "%v"} `, err)
				return
			}

			if out != nil && out.StatusCode != 200 {
				rw.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(rw, `{"error": "Unauthorized"}`)
				return
			}
		}
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
