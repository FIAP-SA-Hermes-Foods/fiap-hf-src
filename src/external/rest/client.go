package rest

import (
	"fiap-hf-src/src/base/interfaces"
	"net/http"
)

var _ interfaces.ClientHandler = (*handlerClient)(nil)

type handlerClient struct {
	controller interfaces.ClientController
}

func NewHandlerClient(controller interfaces.ClientController) *handlerClient {
	return &handlerClient{controller: controller}
}

func (h *handlerClient) Handler(rw http.ResponseWriter, req *http.Request) {

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
