package rest

import (
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fmt"
	"net/http"
)

var _ interfaces.ClientHandler = (*handlerClient)(nil)

type handlerClient struct {
	controller interfaces.ClientController
	userAuth   interfaces.UserAuth
}

func NewHandlerClient(controller interfaces.ClientController, userAuth interfaces.UserAuth) *handlerClient {
	return &handlerClient{controller: controller, userAuth: userAuth}
}

func (h *handlerClient) Handler(rw http.ResponseWriter, req *http.Request) {

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

	var routesClient = map[string]http.HandlerFunc{
		"get hermes_foods/client/{cpf}": h.controller.GetClientByCPF,
		"post hermes_foods/client":      h.controller.SaveClient,
	}

	handler, err := router(req.Method, req.URL.Path, routesClient)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}
