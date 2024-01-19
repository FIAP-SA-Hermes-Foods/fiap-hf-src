package web

import (
	"bytes"
	"encoding/json"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/entity"
	com "fiap-hf-src/internal/core/entity/common"
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
	apiHToken := req.Header.Get("Auth-token")

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	var routeProducts = map[string]http.HandlerFunc{
		"get hermes_foods/product":         h.getProductByCategory,
		"post hermes_foods/product":        h.saveProduct,
		"put hermes_foods/product/{id}":    h.updateProductByID,
		"delete hermes_foods/product/{id}": h.deleteProductByID,
	}

	if err := tokenValidate(apiHToken); err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"error": "not authorized"} `))
		return
	}

	route := ""

	for k := range routeProducts {
		isValidRoute, rr, m := ValidRoute(k, req.URL.Path, req.Method)
		if isValidRoute && m == strings.ToLower(req.Method) {
			route = rr
		}
	}

	if handler, ok := routeProducts[route]; ok {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.URL.Path + ` not found"} `))
}

func (h handlerProduct) saveProduct(rw http.ResponseWriter, req *http.Request) {
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
		Category: com.Category{
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
	id := getID("product", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update product by ID: %v"} `, err)
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
		Category: com.Category{
			Value: reqProduct.Category,
		},
		Image:       reqProduct.Image,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		DeactivatedAt: com.DeactivatedAt{
			Value: nil,
		},
	}

	if len(reqProduct.DeactivatedAt) > 0 {
		product.DeactivatedAt.Value = new(time.Time)
		if err := product.DeactivatedAt.SetTimeFromString(reqProduct.DeactivatedAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to update product: %v"} `, err)
			return
		}
	}

	p, err := h.App.UpdateProductByID(idconv, product)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(p.MarshalString()))
}

func (h handlerProduct) deleteProductByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("product", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if req.Method != http.MethodDelete {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	if err := h.App.DeleteProductByID(idconv); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to delete product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"OK"}`))
}

func (h handlerProduct) getProductByCategory(rw http.ResponseWriter, req *http.Request) {
	category := req.URL.Query().Get("category")

	pList, err := h.App.GetProductByCategory(category)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get product by category: %v"} `, err)
		return
	}

	if pList == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "product not found"}`))
		return
	}

	b, err := json.Marshal(pList)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get product by category: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}
