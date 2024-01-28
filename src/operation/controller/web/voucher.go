package web

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
	"fiap-hf-src/src/operation/presenter"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HandlerVoucher interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerVoucher struct {
	App interfaces.HermesFoodsApp
}

func NewHandlerVoucher(app interfaces.HermesFoodsApp) HandlerVoucher {
	return handlerVoucher{App: app}
}

func (h handlerVoucher) Handler(rw http.ResponseWriter, req *http.Request) {
	apiHToken := req.Header.Get("Auth-token")

	var routeVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/voucher/{id}": h.getVoucherByID,
		"post hermes_foods/voucher":     h.saveVoucher,
		"put hermes_foods/voucher/{id}": h.updateVoucherByID,
	}

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	handler, err := router(req.Method, req.URL.Path, routeVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handlerVoucher) saveVoucher(rw http.ResponseWriter, req *http.Request) {
	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	voucher := entity.Voucher{
		Code:       reqVoucher.Code,
		Porcentage: reqVoucher.Porcentage,
	}

	if len(reqVoucher.ExpiresAt) > 0 {
		voucher.ExpiresAt.Value = new(time.Time)
		if err := voucher.ExpiresAt.SetTimeFromString(reqVoucher.ExpiresAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
			return
		}
	}

	v, err := h.App.SaveVoucher(voucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(presenter.MarshalString(v)))
}

func (h handlerVoucher) updateVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher by ID: %v"} `, err)
		return
	}

	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	voucher := entity.Voucher{
		Code:       reqVoucher.Code,
		Porcentage: reqVoucher.Porcentage,
	}

	if len(reqVoucher.ExpiresAt) > 0 {
		voucher.ExpiresAt.Value = new(time.Time)
		if err := voucher.ExpiresAt.SetTimeFromString(reqVoucher.ExpiresAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
			return
		}
	}

	v, err := h.App.UpdateVoucherByID(idconv, voucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(presenter.MarshalString(v)))
}

func (h handlerVoucher) getVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get voucher by ID: %v"} `, err)
		return
	}

	v, err := h.App.GetVoucherByID(idconv)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(presenter.MarshalString(v)))
}
