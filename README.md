# Run commands in the project root

```shell
go mod init github.com/adminoid/vuego
go run cmd/app/main.go
```

or through make:
```shell
make build # build binary
make test # run tests
make docker # run postgresql docker container
make migrate-up # load migrations into docker container
make migrate-down # unload migrations into docker container
```

for environment edit `.env` file
