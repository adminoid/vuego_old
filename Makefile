include .env
export

.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: check
check:
	echo ${DatabaseURL}

.PHONY: migrate
migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

.DEFAULT_GOAL := build
