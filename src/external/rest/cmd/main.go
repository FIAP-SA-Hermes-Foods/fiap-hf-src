package main

import (
	"bufio"
	"context"
	"fiap-hf-src/src/core/useCase"
	"fiap-hf-src/src/external/db"
	"fiap-hf-src/src/external/db/rds/postgres"
	httpExt "fiap-hf-src/src/external/http"
	"fiap-hf-src/src/external/rest"
	"fiap-hf-src/src/operation/controller/web"
	gatewayDB "fiap-hf-src/src/operation/gateway/db"
	gatewayHTTP "fiap-hf-src/src/operation/gateway/http"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

	dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

	if err != nil {
		log.Fatalf("error: %v", err)

	}

	psql := postgres.NewPostgresDB(
		ctx,
		os.Getenv("DB_REGION"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbDuration,
	)

	defer psql.Close()

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

	paymentAPI := httpExt.NewMercadoPagoAPI(urlAPI, headersAPI, du)
	paymentGateway := gatewayHTTP.NewMercadoPagoAPI(paymentAPI)
	paymentUseCase := useCase.NewPaymentUseCase(paymentGateway)

	clientDB := db.NewClientDB(ctx, psql)
	clientGateway := gatewayDB.NewClientGateway(clientDB)
	clientService := useCase.NewClientUseCase(clientGateway)

	orderDB := db.NewOrderDB(ctx, psql)
	orderGateway := gatewayDB.NewOrderGateway(orderDB)
	orderService := useCase.NewOrderUseCase(orderGateway)

	orderProductDB := db.NewOrderProductDB(ctx, psql)
	orderProductgateway := gatewayDB.NewOrderProductGateway(orderProductDB)
	orderProductService := useCase.NewOrderProductUseCase(orderProductgateway)

	productDB := db.NewProductDB(ctx, psql)
	productgateway := gatewayDB.NewProductGateway(productDB)
	productService := useCase.NewProductUseCase(productgateway)

	voucherDB := db.NewVoucherDB(ctx, psql)
	voucherGateway := gatewayDB.NewVoucherGateway(voucherDB)
	voucherService := useCase.NewVoucherUseCase(voucherGateway)

	app := useCase.NewHermesFoodsApp(
		ctx,
		paymentUseCase,
		clientService,
		orderService,
		orderProductService,
		productService,
		voucherService,
	)

	controllersClient := web.NewClientController(app)
	controllersOrder := web.NewOrderController(app)
	controllersProduct := web.NewProductController(app)
	controllersVoucher := web.NewVoucherController(app)

	configAws := aws.NewConfig()
	configAws.Region = aws.String("us-east-1")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *configAws,
		SharedConfigState: session.SharedConfigEnable,
	}))

	userAuth := httpExt.NewUserAuth(ctx, os.Getenv("USER_AUTH_FUNC_NAME"), *sess)

	handlersClient := rest.NewHandlerClient(controllersClient, userAuth)
	handlersOrder := rest.NewHandlerOrder(controllersOrder, userAuth)
	handlersProduct := rest.NewHandlerProduct(controllersProduct, userAuth)
	hanldersVoucher := rest.NewHandlerVoucher(controllersVoucher, userAuth)

	router.Handle("/hermes_foods/health", http.StripPrefix("/", rest.Middleware(web.HealthCheck)))

	router.Handle("/hermes_foods/client/", http.StripPrefix("/", rest.Middleware(handlersClient.Handler)))
	router.Handle("/hermes_foods/client", http.StripPrefix("/", rest.Middleware(handlersClient.Handler)))

	router.Handle("/hermes_foods/order/", http.StripPrefix("/", rest.Middleware(handlersOrder.Handler)))
	router.Handle("/hermes_foods/order", http.StripPrefix("/", rest.Middleware(handlersOrder.Handler)))

	router.Handle("/hermes_foods/product/", http.StripPrefix("/", rest.Middleware(handlersProduct.Handler)))
	router.Handle("/hermes_foods/product", http.StripPrefix("/", rest.Middleware(handlersProduct.Handler)))

	router.Handle("/hermes_foods/voucher/", http.StripPrefix("/", rest.Middleware(hanldersVoucher.Handler)))
	router.Handle("/hermes_foods/voucher", http.StripPrefix("/", rest.Middleware(hanldersVoucher.Handler)))

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
