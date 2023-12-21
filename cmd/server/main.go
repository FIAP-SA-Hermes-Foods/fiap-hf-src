package main

import (
	"bufio"
	"context"
	cRepo "fiap-hf-src/internal/adapters/driven/repository/client"
	oRepo "fiap-hf-src/internal/adapters/driven/repository/order"
	opRepo "fiap-hf-src/internal/adapters/driven/repository/order_product"
	pRepo "fiap-hf-src/internal/adapters/driven/repository/product"
	vRepo "fiap-hf-src/internal/adapters/driven/repository/voucher"
	apiMercadoPago "fiap-hf-src/internal/adapters/driver/http/api-mercadoPago"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/service"
	"fiap-hf-src/internal/handler/web"
	"fiap-hf-src/pkg/postgres"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

	clientRepo, clientService := cRepo.NewClientRepository(ctx, db), service.NewClientService(nil)
	orderRepo, orderService := oRepo.NewOrderRepository(ctx, db), service.NewOrderService(nil)
	orderProductRepo, orderProductService := opRepo.NewOrderProductRepository(ctx, db), service.NewOrderProductService(nil)
	productRepo, productService := pRepo.NewProductRepository(ctx, db), service.NewProductService(nil)
	voucherRepo, voucherService := vRepo.NewVoucherRepository(ctx, db), service.NewVoucherService(nil)

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

	router.HandleFunc("/hermes_foods/health", web.HealthCheck)
	router.HandleFunc("/hermes_foods/client/", handlersClient.Handler)
	router.HandleFunc("/hermes_foods/order/", handlersOrder.Handler)
	router.HandleFunc("/hermes_foods/product/", handlersProduct.Handler)
	router.HandleFunc("/hermes_foods/voucher/", hanldersVoucher.Handler)

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
		indexComment := strings.Index(sc.Text(), "#")
		if indexComment != -1 && len(strings.TrimSpace(sc.Text()[:indexComment])) == 0 {
			continue
		}
		envEqualSign := strings.Index(sc.Text(), "=")
		if envEqualSign != -1 {
			envMatchKey := sc.Text()[:envEqualSign]
			envMatchValue := sc.Text()[envEqualSign+1:]
			if len(envMatchKey) != 0 || len(envMatchValue) != 0 {
				err := os.Setenv(envMatchKey, envMatchValue)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}
