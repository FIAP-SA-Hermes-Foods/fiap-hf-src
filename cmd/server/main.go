package main

import (
	"context"
	"fiap-hf-src/infrastructure/db/postgres"
	cRepo "fiap-hf-src/internal/adapters/driven/repository/client"
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

	clientRepo, clientService := cRepo.NewClientRepository(ctx, db), service.NewClientService()

	app := application.NewHermesFoodsApp(clientRepo, clientService)
	handlersClient := ui.NewHandlerClient(app)

	router.HandleFunc("/hermes_foods", server)
	router.HandleFunc("/hermes_foods/health", ui.HealthCheck)
	router.HandleFunc("/hermes_foods/client", handlersClient.Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func server(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "Let's kill hunger fast!"}`))
}
