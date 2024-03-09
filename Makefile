.PHONY: run build

run: build
	./main

build: swagger
	go build -o main

swagger: swagfmt
	swag init -g http/apis/*.go

swagfmt:
	swag fmt -g http/apis/*.go

#export PATH=$(go env GOPATH)/bin:$PATH