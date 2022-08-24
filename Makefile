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
	go build main.go
.PHONY:build

run:vet
	go run main.go
.PHONY:run