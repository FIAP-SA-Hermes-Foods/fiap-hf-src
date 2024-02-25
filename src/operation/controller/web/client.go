package web

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"fmt"
	"net/http"
	"strings"
)

var _ interfaces.ClientController = (*clientController)(nil)

type clientController struct {
	app interfaces.HermesFoodsUseCase
}

func NewClientController(app interfaces.HermesFoodsUseCase) *clientController {
	return &clientController{app: app}
}

func (h *clientController) SaveClient(rw http.ResponseWriter, req *http.Request) {
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

	c, err := h.app.SaveClient(reqClient)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save client: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(c)))
}

func (h *clientController) GetClientByCPF(rw http.ResponseWriter, req *http.Request) {
	cpf := getCpf(req.URL.Path)

	c, err := h.app.GetClientByCPF(cpf)

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
	rw.Write([]byte(ps.MarshalString(c)))
}

func getCpf(url string) string {
	indexCpf := strings.Index(url, "client/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], "client/", "")
}
