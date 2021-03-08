.PHONY: help docker build build-go lint lint-go test test-go clean clean-full copy-config post-lint

SHELL=/bin/bash -o pipefail

.DEFAULT_GOAL := help
GO_PATH := $(shell go env GOPATH 2> /dev/null)

help: ## Display general help about this command
	@echo 'Makefile targets:'
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' Makefile \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/    \1 :: \3/p' \
	| column -t -c 1  -s '::'

docker:
	docker build -t fuzzingbits/quiet-hacker-news:latest .

build: build-go ## Build the application

build-go:
	go build -ldflags='-s -w' -o $(CURDIR)/var/quiet-hacker-news .
	@ln -sf $(CURDIR)/var/quiet-hacker-news $(GO_PATH)/bin/quiet-hacker-news

lint: lint-go ## Lint the application

lint-go:
	@cd ; go get golang.org/x/lint/golint
	@cd ; go get golang.org/x/tools/cmd/goimports
	go get -d ./...
	go mod tidy
	gofmt -s -w .
	go vet ./...
	golint -set_exit_status=1 ./...
	goimports -w .

test: test-go ## Test the application

test-go:
	@mkdir -p var/
	@go test -race -cover -coverprofile  var/coverage.txt ./...
	@go tool cover -func var/coverage.txt | awk '/^total/{print $$1 " " $$3}'

clean: ## Remove files listed in .gitignore (possibly with some exceptions)
	git clean -Xdff

clean-full:
	git clean -Xdff

copy-config: ## Copy missing config files into place

post-lint:
	@git diff --exit-code --quiet || (echo 'There should not be any changes after the lint runs' && git status && exit 1;)
