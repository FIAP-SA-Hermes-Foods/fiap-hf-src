package rest

import (
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fmt"
	"net/http"
)

var _ interfaces.OrderHandler = (*handlerOrder)(nil)

type handlerOrder struct {
	controller interfaces.OrderController
	userAuth   interfaces.UserAuth
}

func NewHandlerOrder(controller interfaces.OrderController, userAuth interfaces.UserAuth) *handlerOrder {
	return &handlerOrder{controller: controller, userAuth: userAuth}
}

func (h *handlerOrder) Handler(rw http.ResponseWriter, req *http.Request) {

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
