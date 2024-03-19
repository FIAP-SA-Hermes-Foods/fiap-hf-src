package rest

import (
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fmt"
	"net/http"
)

var _ interfaces.VoucherHandler = (*handlerVoucher)(nil)

type handlerVoucher struct {
	controller interfaces.VoucherController
	userAuth   interfaces.UserAuth
}

func NewHandlerVoucher(controller interfaces.VoucherController, userAuth interfaces.UserAuth) *handlerVoucher {
	return &handlerVoucher{controller: controller, userAuth: userAuth}
}

func (h handlerVoucher) Handler(rw http.ResponseWriter, req *http.Request) {

	userAuthStr := req.Header.Get("user-auth")

	if len(userAuthStr) > 0 {

		var uAuth dto.UserInput

		if err := json.Unmarshal([]byte(userAuthStr), &uAuth); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "%v"} `, err)
			return
		}

		if uAuth.User != nil && uAuth.User.WantRegister {
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
