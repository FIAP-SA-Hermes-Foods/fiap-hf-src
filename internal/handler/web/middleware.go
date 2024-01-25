package web

import (
	"net/http"
)

type Middlewar interface {
	CORS(h http.HandlerFunc) http.HandlerFunc
	CheckRoutes(route, desiredRoute string) error
}

func Middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		rw.Header().Set("Content-Type", "application-json")

		h.ServeHTTP(rw, req)
	}
}
