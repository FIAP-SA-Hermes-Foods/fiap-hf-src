package api_mercadoPago

import (
	"bytes"
	"context"
	"encoding/json"
	"fiap-hf-src/internal/core/entity"
	l "fiap-hf-src/pkg/logger"
	"io"
	"net/http"
	"strings"
	"time"
)

type MercadoPagoAPI interface {
	DoPayment(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error)
}

type clientHTTPRequest interface {
	Do(req *http.Request) (*http.Response, error)
}

var httpClient clientHTTPRequest = &http.Client{}

type mercadoPagoAPI struct {
	URL     string
	Headers map[string]string
	Timeout time.Duration
}

func NewMercadoPagoAPI(url string, headers map[string]string, timeout time.Duration) MercadoPagoAPI {
	return mercadoPagoAPI{URL: url, Headers: headers, Timeout: timeout}
}

func (m mercadoPagoAPI) DoPayment(ctx context.Context, input entity.InputPaymentAPI) (*entity.OutputPaymentAPI, error) {

	ctxReq, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	l.Infof("DoPayment received input: ", " | ", input.MarshalString())

	var buff bytes.Buffer

	if _, err := buff.ReadFrom(strings.NewReader(input.MarshalString())); err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctxReq, http.MethodPost, m.URL, &buff)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	for k, v := range m.Headers {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	cleanSpaces(&rBody)

	l.Infof("received httpAdapter response: ", " | ", string(rBody))

	var out = new(entity.OutputPaymentAPI)

	out.HTTPStatus = resp.StatusCode

	if err := json.Unmarshal(rBody, out); err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	l.Infof("DoPayment output: ", " | ", out.MarshalString())
	return out, nil
}

func cleanSpaces(b *[]byte) {
	*b = bytes.ReplaceAll(*b, []byte("  "), []byte(" "))
	*b = bytes.ReplaceAll(*b, []byte("\t"), []byte(""))
	*b = bytes.ReplaceAll(*b, []byte("\n"), []byte(""))
}
