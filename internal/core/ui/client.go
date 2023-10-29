package ui

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"fmt"
	"net/http"
	"strings"
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
	if strings.ContainsAny("/client/", req.URL.Path) {
		switch req.Method {
		case http.MethodPost:
			h.handlerSaveClient(rw, req)
		case http.MethodGet:
			if len(getCpf(req.URL.Path)) > 0 {
				h.handlerGetClientByCPF(rw, req)
			}
		}
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

	resp := entity.RequestClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(resp.MarshalString()))
}

func (h handlerClient) handlerGetClientByCPF(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	cpf := getCpf(req.URL.Path)

	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

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

	resp := entity.RequestClient{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF.Value,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(),
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(resp.MarshalString()))
}

func getCpf(url string) string {
	indexCpf := strings.Index(url, "client/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], "client/", "")
}
