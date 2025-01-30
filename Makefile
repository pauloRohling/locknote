PACKAGES = ./internal/domain/... ./internal/application/... ./internal/persistence/... ./internal/presentation/...
POSTGRES_URL = "postgres://postgres:postgres@localhost:5432/locknote?sslmode=disable"
MIGRATIONS_PATH = ./db

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## install: install all dependencies
.PHONY: install
install:
	go mod tidy -e

## up: start the docker-compose
.PHONY: up
up:
	docker-compose up -d

## down: stop the docker-compose
.PHONY: down
down:
	docker-compose down

## test: run all tests
.PHONY: test
test:
	go test -race -failfast -buildvcs $(PACKAGES)

## test/c: run all tests and display coverage
.PHONY: test/c
test/c:
	go test -v -race -buildvcs -coverprofile=./tmp/coverage.out $(PACKAGES)
	go tool cover -html=./tmp/coverage.out

## test/b: run all benchmark tests
.PHONY: test/b
test/b:
	go test -bench=. -benchmem $(PACKAGES)

## test/v: run all tests in verbose mode
.PHONY: test/v
test/v:
	go test -v -race -failfast -buildvcs $(PACKAGES)


## run: run the application in watch mode
.PHONY: run
run:
	air -c .air.toml
