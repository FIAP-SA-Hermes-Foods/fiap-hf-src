package main

import (
	"context"
	"fiap-hf-src/infrastructure/db/postgres"
	cRepo "fiap-hf-src/internal/adapters/driven/repository/client"
	oRepo "fiap-hf-src/internal/adapters/driven/repository/order"
	opRepo "fiap-hf-src/internal/adapters/driven/repository/order_product"
	pRepo "fiap-hf-src/internal/adapters/driven/repository/product"
	apiMercadoPago "fiap-hf-src/internal/adapters/driver/http/api-mercadoPago"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/service"
	"fiap-hf-src/internal/core/ui"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

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

	app := application.NewHermesFoodsApp(ctx, paymentApi, clientRepo, orderRepo, orderProductRepo, productRepo, clientService, orderService, orderProductService, productService)
	handlersClient := ui.NewHandlerClient(app)
	handlersOrder := ui.NewHandlerOrder(app)
	handlersProduct := ui.NewHandlerProduct(app)

	router.HandleFunc("/hermes_foods/health", ui.HealthCheck)
	router.HandleFunc("/hermes_foods/client/", handlersClient.Handler)
	router.HandleFunc("/hermes_foods/order/", handlersOrder.Handler)
	router.HandleFunc("/hermes_foods/product/", handlersProduct.Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
