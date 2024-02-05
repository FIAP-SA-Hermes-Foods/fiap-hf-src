package web

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	"fiap-hf-src/src/core/entity"
	ps "fiap-hf-src/src/operation/presenter/strings"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var _ interfaces.VoucherController = (*voucherController)(nil)

type voucherController struct {
	app interfaces.HermesFoodsUseCase
}

func NewVoucherController(app interfaces.HermesFoodsUseCase) *voucherController {
	return &voucherController{app: app}
}

func (h *voucherController) SaveVoucher(rw http.ResponseWriter, req *http.Request) {
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

	reqVoucher.ExpiresAt = voucher.CreatedAt.Format()

	v, err := h.app.SaveVoucher(reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *voucherController) UpdateVoucherByID(rw http.ResponseWriter, req *http.Request) {
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

	reqVoucher.ExpiresAt = voucher.ExpiresAt.Format()

	v, err := h.app.UpdateVoucherByID(idconv, reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *voucherController) GetVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get voucher by ID: %v"} `, err)
		return
	}

	v, err := h.app.GetVoucherByID(idconv)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(v)))
}
