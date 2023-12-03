package ui

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/domain/entity"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HandlerVoucher interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerVoucher struct {
	App application.HermesFoodsApp
}

func NewHandlerVoucher(app application.HermesFoodsApp) HandlerVoucher {
	return handlerVoucher{App: app}
}

func (h handlerVoucher) Handler(rw http.ResponseWriter, req *http.Request) {
	apiHToken := req.Header.Get("Auth-token")

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	rw.Header().Add("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	rw.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
	rw.Header().Add("Content-Type", "application-json")

	if strings.ContainsAny("/voucher/", req.URL.Path) {
		switch req.Method {
		case http.MethodPost:
			h.saveVoucher(rw, req)
		case http.MethodPut:
			if len(getID("voucher", req.URL.Path)) > 0 {
				h.updateVoucherByID(rw, req)
			}
		case http.MethodGet:
			h.getVoucherByID(rw, req)
		default:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`{"error": "route not found"} `))
			return
		}
	}
}

func (h handlerVoucher) saveVoucher(rw http.ResponseWriter, req *http.Request) {
	var buff bytes.Buffer

	var reqVoucher entity.RequestVoucher

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
	rw.Write([]byte(v.MarshalString()))
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

	var reqVoucher entity.RequestVoucher

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

	p, err := h.App.UpdateVoucherByID(idconv, voucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(p.MarshalString()))
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
	rw.Write([]byte(v.MarshalString()))
}
