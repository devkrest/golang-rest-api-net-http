# Binary name
BINARY_NAME=Golang-api

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

.PHONY: all build clean test run docs help

all: help

help:
	@echo "Golang API - Master Makefile"
	@echo "Usage:"
	@echo "  make run    - Run the application with Air (live reload)"
	@echo "  make build  - Build the application binary"
	@echo "  make docs   - Generate Swagger documentation"
	@echo "  make tidy   - Clean and update Go modules"
	@echo "  make test   - Run all tests"
	@echo "  make clean  - Remove binary files"

run:
	air

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/api

docs:
	swag init -g cmd/main.go

tidy:
	$(GOMOD) tidy
	$(GOMOD) vendor

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf bin/
	rm -rf vendor/
