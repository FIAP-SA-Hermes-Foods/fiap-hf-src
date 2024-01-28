package web

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/operation/presenter"
	"fmt"
	"net/http"
	"strings"
)

type HandlerClient interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerClient struct {
	App interfaces.HermesFoodsApp
}

func NewHandlerClient(app interfaces.HermesFoodsApp) HandlerClient {
	return handlerClient{App: app}
}

func (h handlerClient) Handler(rw http.ResponseWriter, req *http.Request) {
	apiHToken := req.Header.Get("Auth-token")

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	var routesClient = map[string]http.HandlerFunc{
		"get hermes_foods/client/{cpf}": h.handlerGetClientByCPF,
		"post hermes_foods/client":      h.handlerSaveClient,
	}

	handler, err := router(req.Method, req.URL.Path, routesClient)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handlerClient) handlerSaveClient(rw http.ResponseWriter, req *http.Request) {
	var (
		buff      bytes.Buffer
		reqClient dto.RequestClient
	)

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqClient); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	client := reqClient.Client()

	c, err := h.App.SaveClient(client)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save client: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(presenter.MarshalString(c)))
}

func (h handlerClient) handlerGetClientByCPF(rw http.ResponseWriter, req *http.Request) {
	cpf := getCpf(req.URL.Path)

	c, err := h.App.GetClientByCPF(cpf)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get client by ID: %v"} `, err)
		return
	}

	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "client not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(presenter.MarshalString(c)))
}

func getCpf(url string) string {
	indexCpf := strings.Index(url, "client/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], "client/", "")
}
