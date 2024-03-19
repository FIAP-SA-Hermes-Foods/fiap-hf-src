package api

import (
	"encoding/json"
	"fiap-hf-src/src/base/dto"
	"log"
	"net/http"
)

// Mock HTTP request
type httpMock struct {
	WantOut *http.Response
	WantErr error
}

func (h httpMock) Do(*http.Request) (*http.Response, error) {
	return h.WantOut, h.WantErr
}

// Mock input to DoPaymentFunction
var (
	successInput = `{"price":35.5,"client":{"name":"SomeoneTest","cpf":{"Value":"00000011111"},"email":"someone@email.com","phone":"510000119999"}}`
)

func mockInputDoPayment(strIn string) (in dto.InputPaymentAPI) {
	if err := json.Unmarshal([]byte(strIn), &in); err != nil {
		log.Printf("error in mock input: %v", err)
	}
	return in
}

// Mock output to DoPaymentFunction
var (
	successResponseAPIMock = `{"paymentStatus":"Paid","httpStatus":200}`
)

func mockOutputDoPayment(strIn string) (out dto.OutputPaymentAPI) {
	if err := json.Unmarshal([]byte(strIn), &out); err != nil {
		log.Printf("error in mock input: %v", err)
	}
	return out
}

var (
	successDoPaymentOutputMock = `{"paymentStatus":"Paid","httpStatus":200}`
	errorDoPaymentOutputMock   = `{"httpStatus": 500,"error": {"message": "payment method not accepted by Mercado Pago", "code": "F4008"}}`
)
