package web

import (
	"encoding/base64"
	l "fiap-hf-src/pkg/logger"
	"fmt"
	"os"
	"strings"
)

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
	indexCpf := strings.Index(url, handlerName+"/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], handlerName+"/", "")
}
