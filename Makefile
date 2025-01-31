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

.PHONY: migration
## migration name=?: create a new migration
migration:
	docker run --rm -u 1000:1000 -v .:/migrations --network host migrate/migrate create -ext=sql -dir=/migrations/$(MIGRATIONS_PATH) -seq $(name)

.PHONY: migrate-up
## migrate-up: execute all migrations
migrate-up:
	docker run --rm -u 1000:1000 -v .:/migrations --network host migrate/migrate -verbose -path=/migrations/$(MIGRATIONS_PATH) -database $(POSTGRES_URL) up

.PHONY: migrate-down
## migrate-down: revert all migrations
migrate-down:
	docker run --rm -u 1000:1000 -v .:/migrations --network host migrate/migrate -verbose -path=/migrations/$(MIGRATIONS_PATH) -database $(POSTGRES_URL) down -all

.PHONY: sql
## sql: generate sql code
sql:
	docker run --rm -u 1000:1000 -v .:/src -w /src sqlc/sqlc generate -f="./sqlc.yml"

## mock: generate mocks
.PHONY: mock
mock:
	docker run --rm -u 1000:1000 -v .:/src -w /src vektra/mockery:v2.51 --all

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
