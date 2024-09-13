.PHONY: lint test 

default: lint test

lint:
	golangci-lint run

test:
	go test -v -cover ./...
