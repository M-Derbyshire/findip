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

create-snapshot-directories:
	mkdir -p snapshots/ip

test: create-snapshot-directories
	./scripts/unit_tests.sh
.PHONY:test

e2e-test:
	./scripts/e2e_tests.sh
.PHONY:e2e-test

update-test-snapshots: create-snapshot-directories
	go test ./ip... -v -update
.PHONY:update-test-snapshots