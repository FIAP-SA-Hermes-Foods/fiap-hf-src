package web

import (
	"strconv"
	"strings"
)

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
