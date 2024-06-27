.PHONY:
.SILENT:

init:
	go mod download

gorun:
	go run cmd/main.go

gobuild:
	go build cmd/main.go

linter:
	golangci-lint run

build:
	docker-compose build

run:
	docker compose up -d

stop:
	docker compose down
