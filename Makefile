.DEFAULT_GOAL:=run

fmt:
	go fmt ./...
.PHONY:fmt

lint:fmt
	golint ./...
.PHONY:lint

vet:fmt
	go vet ./...
	shadow ./...
.PHONY:vet

build:vet
	go build hello.go
.PHONY:build

run:vet
	go run hello.go
.PHONY:run