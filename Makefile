run-local:
	go mod init fiap-hf-src;
	go mod tidy;
	go build -ldflags "-w -s" -o bin/hermesfoods src/external/cmd/server/*.go;
	./bin/hermesfoods;

run-build:	
	./infrastructure/scripts/docker-network.sh;
	docker-compose up --build --force-recreate;

run-build-d:	
	./infrastructure/scripts/docker-network.sh;
	docker-compose up --build --force-recreate -d;

run:
	./infrastructure/scripts/docker-network.sh;
	docker-compose up;

migrate:
	docker exec -it go-hermes-foods-app bash -c 'go build -ldflags "-w -s" -o bin/hermesfoods-migration cmd/migration/*.go ; ./bin/hermesfoods-migration';

tests:
	@docker exec -it go-hermes-foods-app /fiap-hf-src/src/app/infrastructure/scripts/coverage.sh;
	

