package web

import (
	"net/http"
	"strconv"
	"strings"
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

func ValidRoute(route, requestRoute, method string) (bool, string, string) {

	requestRoute = strings.ToLower(method) + " " + requestRoute

	isValid := false

	if requestRoute[len(requestRoute)-1] == '/' {
		requestRoute = requestRoute[:len(requestRoute)-1]
	}

	routeItems := strings.Split(route, "/")
	desiredRouteItems := strings.Split(requestRoute, "/")
	if len(routeItems) != len(desiredRouteItems) {
		return false, "", ""
	}

	idParamVal := ""

	for i := 0; i < len(routeItems); i++ {
		if strings.Contains(routeItems[i], "{") {
			idParamVal = desiredRouteItems[i]
		}
		if routeItems[i] == desiredRouteItems[i] {
			isValid = true
			continue
		}
		isValid = false
	}

	if idParamVal != "" {
		if _, err := strconv.Atoi(idParamVal); err != nil {
			return false, "", ""
		} else {
			isValid = true
		}
	}

	var methodReturn string
	if len(route) > 0 {
		methodReturn = strings.Split(route, " ")[0]
	}

	return isValid, route, methodReturn
}

func cleanMethods(url *string) {
	if url == nil {
		return
	}

	var urlVal = *url

	if strings.Contains(*url, "get") {
		urlVal = urlVal[3:]
	} else if strings.Contains(*url, "post") {
		urlVal = urlVal[4:]
	} else if strings.Contains(*url, "patch") {
		urlVal = urlVal[5:]
	} else if strings.Contains(*url, "put") {
		urlVal = urlVal[3:]
	} else if strings.Contains(*url, "delete") {
		urlVal = urlVal[6:]
	}

	*url = urlVal
}
