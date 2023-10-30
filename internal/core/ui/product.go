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
	"time"
)

type HandlerProduct interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}

type handlerProduct struct {
	App application.HermesFoodsApp
}

func NewHandlerProduct(app application.HermesFoodsApp) HandlerProduct {
	return handlerProduct{App: app}
}

func (h handlerProduct) Handler(rw http.ResponseWriter, req *http.Request) {
	if strings.ContainsAny("/product/", req.URL.Path) {
		switch req.Method {
		case http.MethodPost:
			h.saveProduct(rw, req)
		case http.MethodPut:
			if len(getID("product", req.URL.Path)) > 0 {
				h.updateProductByID(rw, req)
			}

		case http.MethodGet:
		case http.MethodDelete:
		default:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`{"error": "route not found"} `))
			return
		}
	}
}

func (h handlerProduct) saveProduct(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqProduct entity.RequestProduct

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqProduct); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	product := entity.Product{
		Name: reqProduct.Name,
		Category: valueObject.Category{
			Value: reqProduct.Category,
		},
		Image:       reqProduct.Image,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
	}

	p, err := h.App.SaveProduct(product)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(p.MarshalString()))
}

func (h handlerProduct) updateProductByID(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getID("product", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if req.Method != http.MethodPut {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqProduct entity.RequestProduct

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqProduct); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	product := entity.Product{
		Name: reqProduct.Name,
		Category: valueObject.Category{
			Value: reqProduct.Category,
		},
		Image:       reqProduct.Image,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		DeactivatedAt: valueObject.DeactivatedAt{
			Value: nil,
		},
	}

	if len(reqProduct.DeactivatedAt) > 0 {
		product.DeactivatedAt.Value = new(time.Time)
		product.DeactivatedAt.SetTimeFromString(reqProduct.DeactivatedAt)
	}

	p, err := h.App.UpdateProductByID(idconv, product)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(p.MarshalString()))
}
