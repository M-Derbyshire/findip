.DEFAULT_GOAL := build-dev

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build-dev: vet lint
	go build .
.PHONY:build-dev

build-prod: vet lint
	go build -ldflags=-w .
.PHONY:build-prod

test:
	go test -v ./...
.PHONY:test
