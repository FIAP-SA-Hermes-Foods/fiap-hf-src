package main

import (
	"bufio"
	"context"
	"fiap-hf-src/src/core/useCase/application"
	useCaseClient "fiap-hf-src/src/core/useCase/client"
	useCaseOrder "fiap-hf-src/src/core/useCase/order"
	useCaseOrderProduct "fiap-hf-src/src/core/useCase/order_product"
	useCaseProduct "fiap-hf-src/src/core/useCase/product"
	useCaseVoucher "fiap-hf-src/src/core/useCase/voucher"
	"fiap-hf-src/src/external/db/postgres"
	"fiap-hf-src/src/operation/controller/web"
	apiMercadoPago "fiap-hf-src/src/operation/gateway/http/api-mercadoPago"
	cRepo "fiap-hf-src/src/operation/gateway/repository/client"
	oRepo "fiap-hf-src/src/operation/gateway/repository/order"
	opRepo "fiap-hf-src/src/operation/gateway/repository/order_product"
	pRepo "fiap-hf-src/src/operation/gateway/repository/product"
	vRepo "fiap-hf-src/src/operation/gateway/repository/voucher"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	notFindIndex       = -1
	indexMatchedValue  = 1
	nullMatch          = 0
	commentCharacter   = "#"
	separatorCharacter = "="
)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Fatalf("Error to load .env -> %v", err)
	}
}

func main() {
	printArt()
	go APIMercadoPago()

	router := http.NewServeMux()

	ctx := context.Background()

	db := postgres.NewPostgresDB(
		ctx,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	defer db.Close()

	urlAPI := fmt.Sprintf("http://%s:%s/%s",
		os.Getenv("MERCADO_PAGO_API_HOST"),
		os.Getenv("MERCADO_PAGO_API_PORT"),
		os.Getenv("MERCADO_PAGO_API_URI"),
	)

	headersAPI := map[string]string{
		"Content-type": "application/json",
	}

	du, err := time.ParseDuration(os.Getenv("MERCADO_PAGO_API_TIMEOUT"))

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	paymentApi := apiMercadoPago.NewMercadoPagoAPI(urlAPI, headersAPI, du)

	clientRepo, clientService := cRepo.NewClientRepository(ctx, db), useCaseClient.NewClientService(nil)
	orderRepo, orderService := oRepo.NewOrderRepository(ctx, db), useCaseOrder.NewOrderService(nil)
	orderProductRepo, orderProductService := opRepo.NewOrderProductRepository(ctx, db), useCaseOrderProduct.NewOrderProductService(nil)
	productRepo, productService := pRepo.NewProductRepository(ctx, db), useCaseProduct.NewProductService(nil)
	voucherRepo, voucherService := vRepo.NewVoucherRepository(ctx, db), useCaseVoucher.NewVoucherService(nil)

	app := application.NewHermesFoodsApp(
		ctx,
		paymentApi,
		clientRepo,
		orderRepo,
		orderProductRepo,
		productRepo,
		voucherRepo,
		clientService,
		orderService,
		orderProductService,
		productService,
		voucherService,
	)

	handlersClient := web.NewHandlerClient(app)
	handlersOrder := web.NewHandlerOrder(app)
	handlersProduct := web.NewHandlerProduct(app)
	hanldersVoucher := web.NewHandlerVoucher(app)

	router.Handle("/hermes_foods/health", http.StripPrefix("/", web.Middleware(web.HealthCheck)))

	router.Handle("/hermes_foods/client/", http.StripPrefix("/", web.Middleware(handlersClient.Handler)))
	router.Handle("/hermes_foods/client", http.StripPrefix("/", web.Middleware(handlersClient.Handler)))

	router.Handle("/hermes_foods/order/", http.StripPrefix("/", web.Middleware(handlersOrder.Handler)))
	router.Handle("/hermes_foods/order", http.StripPrefix("/", web.Middleware(handlersOrder.Handler)))

	router.Handle("/hermes_foods/product/", http.StripPrefix("/", web.Middleware(handlersProduct.Handler)))
	router.Handle("/hermes_foods/product", http.StripPrefix("/", web.Middleware(handlersProduct.Handler)))

	router.Handle("/hermes_foods/voucher/", http.StripPrefix("/", web.Middleware(hanldersVoucher.Handler)))
	router.Handle("/hermes_foods/voucher", http.StripPrefix("/", web.Middleware(hanldersVoucher.Handler)))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func defineEnvs(filename string) error {
	file, err := os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error on close -> %v", err)
		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		indexComment := strings.Index(sc.Text(), commentCharacter)
		if indexComment != notFindIndex && len(strings.TrimSpace(sc.Text()[:indexComment])) == nullMatch {
			continue
		}
		envEqualSign := strings.Index(sc.Text(), separatorCharacter)
		if envEqualSign != notFindIndex {
			envMatchKey := sc.Text()[:envEqualSign]
			envMatchValue := sc.Text()[envEqualSign+indexMatchedValue:]
			if len(envMatchKey) != nullMatch || len(envMatchValue) != nullMatch {
				err := os.Setenv(envMatchKey, envMatchValue)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}