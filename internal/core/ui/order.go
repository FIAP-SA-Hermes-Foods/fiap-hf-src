package ui

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
	"fmt"
	"net/http"
	"strconv"
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
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	rw.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
	rw.Header().Add("Content-Type", "application/json")

	if strings.ContainsAny("/order/", req.URL.Path) {
		switch req.Method {
		case http.MethodPost:
			h.saveOrder(rw, req)
		case http.MethodGet:
			if len(getID("order", req.URL.Path)) > 0 {
				h.handlerGetOrderByID(rw, req)
				return
			}
			h.handlerGetOrders(rw, req)
		case http.MethodPatch:
			if len(getID("order", req.URL.Path)) > 0 {
				h.handlerUpdateOrderByID(rw, req)
			}
		default:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`{"error": "route not found"} `))
			return
		}
	}
}

func (h handlerOrder) saveOrder(rw http.ResponseWriter, req *http.Request) {
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
		Items:     reqOrder.Items,
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

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(o.MarshalString()))
}

func (h handlerOrder) handlerGetOrderByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("order", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	o, err := h.App.GetOrderByID(idconv)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if o == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(o.MarshalString()))
}

func (h handlerOrder) handlerGetOrders(rw http.ResponseWriter, req *http.Request) {
	oList, err := h.App.GetOrders()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if oList == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	b, err := json.Marshal(oList)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func (h handlerOrder) handlerUpdateOrderByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("order", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
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

	o, err := h.App.UpdateOrderByID(idconv, order)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if o == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(o.MarshalString()))
}
