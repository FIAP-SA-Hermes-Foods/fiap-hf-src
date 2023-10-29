package main

import (
	"context"
	"fiap-hf-src/infrastructure/db/postgres"
	cRepo "fiap-hf-src/internal/adapters/driven/repository/client"
	oRepo "fiap-hf-src/internal/adapters/driven/repository/order"
	"fiap-hf-src/internal/core/application"
	"fiap-hf-src/internal/core/service"
	"fiap-hf-src/internal/core/ui"
	"log"
	"net/http"
	"os"
)

func main() {
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

	clientRepo, clientService := cRepo.NewClientRepository(ctx, db), service.NewClientService(nil)
	orderRepo, orderService := oRepo.NewOrderRepository(ctx, db), service.NewOrderService(nil)

	app := application.NewHermesFoodsApp(clientRepo, orderRepo, clientService, orderService)
	handlersClient := ui.NewHandlerClient(app)
	handlersOrder := ui.NewHandlerOrder(app)

	router.HandleFunc("/hermes_foods/health", ui.HealthCheck)
	router.HandleFunc("/hermes_foods/client/", handlersClient.Handler)
	router.HandleFunc("/hermes_foods/order/", handlersOrder.Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
