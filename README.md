# Vuego

## Universal boilerplate for nested set postgres structure of site pages

**Based on Golang and Vue.js**

## Run commands in the project root

```shell
go mod init github.com/adminoid/vuego
go run cmd/app/main.go
```

Makefile commands:
```shell
make build # build binary
make test # run tests
make docker # run postgresql docker container
make migrate-up # load migrations into docker container
make migrate-down # unload migrations into docker container
```

for environment edit `.env` file

## Migrations
Created by the following command: `migrate create -ext sql -dir ./migrations -seq init`
If migrations throw error, clean specific table: `TRUNCATE TABLE schema_migrations`

## Hashing passwords
Algorithm: <https://github.com/alexedwards/argon2id>
