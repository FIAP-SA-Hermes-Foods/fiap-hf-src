package main

import (
	"log"
	"net/http"
)

func APIMercadoPago() {
	sMux := http.NewServeMux()
	sMux.HandleFunc("/mercado_pago_api", handlerMercadoPago)
	log.Fatal(http.ListenAndServe(":8081", sMux))
}

func handlerMercadoPago(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"paymentStatus":"Paid","httpStatus":200}`))
}
