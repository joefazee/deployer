.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## run: run the cmd/api application
.PHONY: run
run:
	go run ./cmd/api



current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build: build the cmd/api application
.PHONY: build
build:
	go mod verify
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api



## format: the project using go fmt
.PHONY: format
format:
	go fmt ./...

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## test: run all the tests with covet
.PHONY: test
test:
	@GOFLAGS="-count=1" go test -v -cover -race -vet=off ./...

DB_DSN=postgresql://root:secret@localhost:5455/deployer?sslmode=disable

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	make test
	go mod verify


MIGRATION_PATH=./pkg/db/migrations
## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
	@migrate create -seq -ext=.sql -dir=${MIGRATION_PATH} ${name}


## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up:
	@migrate -path=${MIGRATION_PATH} -database="${DB_DSN}" up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	@migrate -path=${MIGRATION_PATH} -database="${DB_DSN}" down
