run-local:
	go mod init fiap-hf-src;
	go mod tidy;
	go build -ldflags "-w -s" -o bin/hermesfoods cmd/server/*.go;
	./bin/hermesfoods;

run-build:	
	./infrastructure/scripts/docker-network.sh;
	docker-compose up --build --force-recreate;

run:
	./infrastructure/scripts/docker-network.sh;
	docker-compose up;

migration:
	./infrastructure/scripts/docker-migration.sh;
	docker-compose up;

tests:
	@./infrastructure/scripts/coverage.sh
