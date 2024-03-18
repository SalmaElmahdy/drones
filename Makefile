.PHONY: run build

run: build
	./main

build: swagger
	go build -o main

swagger: swagfmt
	swag init -g http/apis/*.go

swagfmt:
	swag fmt -g http/apis/*.go

make_migrations:
	@if [ -z "$(NAME)" ]; then \
        echo "Error: NAME for that migration should set Example:NAME="init""; \
        exit 1; \
    fi
	atlas migrate diff "$(NAME)" --env gorm
#export PATH=$(go env GOPATH)/bin:$PATH

migrate:
	atlas migrate apply --env local --allow-dirty

test:
	@go test -v ./...