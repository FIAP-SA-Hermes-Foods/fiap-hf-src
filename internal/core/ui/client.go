package ui

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"fmt"
	"net/http"
)

type HandlerClient interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerClient struct {
	App application.HermesFoodsApp
}

func NewHandlerClient(app application.HermesFoodsApp) HandlerClient {
	return handlerClient{App: app}
}

func (h handlerClient) Handler(rw http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost:
		h.handlerSaveClient(rw, req)
		return
	case http.MethodGet:
	}
}

func (h handlerClient) handlerSaveClient(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqClient entity.RequestClient

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

	client := entity.Client{
		Name: reqClient.Name,
		CPF: valueObject.Cpf{
			Value: reqClient.CPF,
		},
		Email: reqClient.Email,
	}

	c, err := h.App.SaveClient(client)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save client: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(c.MarshalString()))
}
