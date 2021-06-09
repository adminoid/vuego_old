include .env
export

.DEFAULT_GOAL := build
.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: docker
docker:
	cd build/app && docker compose --env-file .../.env up -d postgres

.PHONY: migrate-up migrate-down
db := postgres://${DbUser}:${DbPwd}@localhost:5432/${DbName}?sslmode=disable
migrate-up:
	migrate -path migrations -database ${db} up
migrate-down:
	migrate -path migrations -database ${db} down
