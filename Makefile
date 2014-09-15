.PHONY: all clean coverage dependencies fmt install lint test

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: clean install test

clean:
	go clean -i ./...
coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
dependencies:
	go get -t -u -v ./...
	go build -v ./...
fmt:
	gofmt -l -w $(ROOT_DIR)/
install: clean
	go install -v ./...
lint: install
	go tool vet -all=true $(ROOT_DIR)/
	golint $(ROOT_DIR)/
test:
	go test ./...
