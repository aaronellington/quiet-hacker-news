.PHONY: full build build-go test test-go lint lint-go fix fix-go watch watch-go clean docs-go docker docker-publish

SHELL=/bin/bash -o pipefail
$(shell git config core.hooksPath ops/git-hooks)
GO_PATH := $(shell go env GOPATH 2> /dev/null)
PATH := /usr/local/bin:$(GO_PATH)/bin:$(PATH)

full: clean lint test build

## Build the project
build: build-go

build-go:
	go generate
	go build -ldflags='-s -w' -o var/build .
	go install .

## Test the project
test: test-go

test-go:
	@mkdir -p var/coverage/go/
	@go install github.com/boumenot/gocover-cobertura@latest
	go test -p 1 -count=1 -cover -coverprofile var/coverage/go/profile.txt ./...
	@go tool cover -func var/coverage/go/profile.txt | awk '/^total/{print $$1 " " $$3}'
	@go tool cover -html var/coverage/go/profile.txt -o var/coverage/go/coverage.html
	@gocover-cobertura < var/coverage/go/profile.txt > var/coverage/go/cobertura-coverage.xml

## Lint the project
lint: lint-go

lint-go:
	go get -d ./...
	go mod tidy
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55
	golangci-lint run ./...

## Fix the project
fix: fix-go

fix-go:
	go mod tidy
	gofmt -s -w .

## Watch the project
watch:
	make -j1 watch-go

watch-go:
	@go install github.com/mitranim/gow@latest
	gow -e=go,html run . | jq

## Clean the project
clean:
	git clean -Xdff --exclude="!.env*local"

## Run the docs server for the project
docs-go:
	@go install golang.org/x/tools/cmd/godoc@latest
	@echo "listening on http://127.0.0.1:6060/pkg/github.com/aaronellington/quiet-hacker-news"
	@godoc -http=127.0.0.1:6060

## Build the Docker image
docker: clean
	docker build -t ghcr.io/aaronellington/quiet-hacker-news:latest .

## Build the Docker image
docker-publish: clean docker
	docker push ghcr.io/aaronellington/quiet-hacker-news:latest
