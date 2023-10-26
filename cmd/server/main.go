package main

import (
	"log"
	"net/http"
)

func main() {
	sMux := http.NewServeMux()
	sMux.HandleFunc("/hermes_foods", server)
	log.Fatal(http.ListenAndServe(":8080", sMux))
}

func server(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "Let's kill hunger fast!"}`))
}
