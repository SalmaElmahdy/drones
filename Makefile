.PHONY: run build

run: build
	./main

build: swagger
	go build -o main

swagger: swagfmt
	swag init -g http/apis/drone.go

swagfmt:
	swag fmt -g http/apis/drone.go
