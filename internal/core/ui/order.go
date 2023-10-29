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

type HandlerOrder interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerOrder struct {
	App application.HermesFoodsApp
}

func NewHandlerOrder(app application.HermesFoodsApp) HandlerOrder {
	return handlerOrder{App: app}
}

func (h handlerOrder) Handler(rw http.ResponseWriter, req *http.Request) {
	if strings.ContainsAny("/order/", req.URL.Path) {
		switch req.Method {
		case http.MethodPost:
			h.saveOrder(rw, req)
		case http.MethodGet:
		}
	}
}

func (h handlerOrder) saveOrder(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqOrder entity.RequestOrder

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqOrder); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	order := entity.Order{
		ClientID:  reqOrder.ClientID,
		VoucherID: reqOrder.VoucherID,
		Status: valueObject.Status{
			Value: reqOrder.Status,
		},
	}

	o, err := h.App.SaveOrder(order)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save order: %v"} `, err)
		return
	}

	resp := entity.RequestOrder{
		ID:               o.ID,
		ClientID:         o.ClientID,
		VoucherID:        o.VoucherID,
		Status:           o.Status.Value,
		VerificationCode: o.VerificationCode.Value,
		CreatedAt:        o.CreatedAt.Format(),
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(resp.MarshalString()))
}
