.PHONY: full build build-go test test-go lint lint-go fix fix-go watch clean docker

SHELL=/bin/bash -o pipefail
GO_PATH := $(shell go env GOPATH 2> /dev/null)
PATH := /usr/local/bin:$(GO_PATH)/bin:$(PATH)

full: clean lint test build

## Build the project
build: build-go

build-go:
	go generate
	go build -ldflags='-s -w' -o var/build .
	@go install .

## Test the project
test: test-go

test-go:
	@mkdir -p var/
	@go test -race -cover -coverprofile  var/coverage.txt ./...
	@go tool cover -func var/coverage.txt | awk '/^total/{print $$1 " " $$3}'

## Lint the project
lint: lint-go

lint-go:
	@go install golang.org/x/lint/golint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	go get -d ./...
	gofmt -s -w .
	go vet ./...
	golint -set_exit_status=1 ./...
	goimports -w .

## Fix the project
fix: fix-go

fix-go:
	go mod tidy
	gofmt -s -w .
	goimports -w .

## Watch the project
watch:

## Clean the project
clean:
	git clean -Xdff

## Build the docker image
docker: clean
	docker build -t aaronellington/quiet-hacker-news .
