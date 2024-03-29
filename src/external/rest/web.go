package rest

import (
	"encoding/base64"
	"errors"
	l "fiap-hf-src/src/base/logger"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func router(reqMethod, path string, routesMap map[string]http.HandlerFunc) (http.HandlerFunc, error) {
	route := ""

	for k := range routesMap {
		isValidRoute, rr, m := validRoute(k, path, reqMethod)
		if isValidRoute && m == strings.ToLower(reqMethod) {
			route = rr
		}
	}

	if handler, ok := routesMap[route]; ok {
		return handler, nil
	}

	return nil, errors.New("route not found")
}

func validRoute(route, requestRoute, method string) (bool, string, string) {

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

func tokenValidate(tokenInput string) error {
	apiHTokenDecode, err := base64.StdEncoding.DecodeString(tokenInput)

	if err != nil {
		l.Warningf("request blocked, invalid token: ", " | ", tokenInput)
		return err
	}

	envToken := os.Getenv("API_TOKEN")

	envTokenDecode, err := base64.StdEncoding.DecodeString(envToken)

	if err != nil {
		l.Errorf("tokenValidate error: ", " | ", err)
		return err
	}

	apiHTokenStr := string(apiHTokenDecode)
	envTokenStr := string(envTokenDecode)

	if len(apiHTokenStr) == 0 || apiHTokenStr != envTokenStr {
		l.Warningf("request blocked, invalid token:", apiHTokenStr)
		return fmt.Errorf("request blocked, invalid token: %s", apiHTokenStr)
	}

	return nil
}

func getID(handlerName, url string) string {
	index := strings.Index(url, handlerName+"/")

	if index == -1 {
		return ""
	}

	id := strings.ReplaceAll(url[index:], handlerName+"/", "")

	if _, err := strconv.Atoi(id); err != nil {
		return ""
	}

	return id
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
