include .env
.PHONY: lint test 

default: lint test

run-tester:
	@KAIROS_USER=$(KAIROS_USER) KAIROS_PASSWORD=$(KAIROS_PASSWORD) KAIROS_IP=$(KAIROS_IP) KAIROS_PORT=$(KAIROS_PORT) \
	go run examples/tester/main.go

lint:
	golangci-lint run

test:
	go test -v -cover ./...

