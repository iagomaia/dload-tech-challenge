.PHONY: all build clean
.PHONY: all vendor clean

include .env
export

db-local:
ifeq ($(shell docker container ps -a --format '{{.Names}}' | grep clinik-db), clinik-db)
	docker start clinik-db
else
	docker run --name clinik-db \
		-e MONGO_INITDB_ROOT_USERNAME=admin \
		-e MONGO_INITDB_ROOT_PASSWORD=mongopw \
		-p $(DB_PORT):27017 -d mongo
endif

ensure-dependencies:
	go mod tidy

test:
	go clean -testcache
	go test ./...

test-cover:
	go clean -testcache
	go test -coverprofile=./cov-report/coverage.out ./...
	go tool cover -html=./cov-report/coverage.out
	
vendor: 
	go mod vendor

run-api: ensure-dependencies vendor
	clear
	go run cmd/api/main.go

build:
ifeq (${OS},(Windows_NT))
	if (!(Test-Path ./build)) { mkdir build }
else
	mkdir -p build
endif
	go build -mod=vendor -v -o ./build ./...
	cp .env ./build/.env
	mkdir -p ./build/static
	cp ./docs/swagger.yaml ./build/static/swagger.yaml

docker-build: build
	VERSION_IMAGE=""
ifdef VERSION_IMAGE
	docker image build --platform=linux/arm64 . -t clinik-school-api:${VERSION_IMAGE}
else
	docker image build --platform=linux/arm64 . -t clinik-school-api
endif

docker-run:
	docker run -p 3000:3000 --env-file .env clinik-school-api