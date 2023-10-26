run-build:
	./infrastructure/scripts/docker-netowrk.sh;
	docker-compose up --build --force-recreate;

run:
	./infrastructure/scripts/docker-netowrk.sh;
	docker-compose up;

migration:
	@go build -ldflags "-w -s" -o bin/hermesfoods-migration cmd/migration/*.go;
	./bin/hermesfoods-migration;

tests:
	@./infrastructure/scripts/coverage.sh
