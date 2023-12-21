package web

import (
	"encoding/base64"
	l "fiap-hf-src/pkg/logger"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func HealthCheck(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	rw.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
	rw.Header().Add("Content-Type", "application-json")

	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func tokenValidate(tokenInput string) error {
	apiHTokenDecode, err := base64.StdEncoding.DecodeString(tokenInput)
	log.Printf("inputheader token -> %v\n", tokenInput)

	if err != nil {
		l.Warningf("request blocked, invalid token: ", " | ", tokenInput)
		return err
	}

	envToken := os.Getenv("API_TOKEN")
	log.Printf("env token -> %v\n", envToken)

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
	indexCpf := strings.Index(url, handlerName+"/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], handlerName+"/", "")
}
