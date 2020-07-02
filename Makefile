-include .env
export

.PHONY: help
GO_PATH := $(shell go env GOPATH)
MODULE := $(shell awk '/^module/ {print $$2}' go.mod)
NAMESPACE := $(shell awk -F "/" '/^module/ {print $$(NF-1)}' go.mod)
PROJECT_NAME := $(shell awk -F "/" '/^module/ {print $$(NF)}' go.mod)

help:
	@echo "Makefile targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' Makefile \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\t\1 :: \3/p' \
	| column -t -c 1  -s '::'

release: githooks clean lint test build ## Do a full clean build

build: ## Build the project
	go build -o $(CURDIR)/var/$(PROJECT_NAME)
	@ln -sf $(CURDIR)/var/$(PROJECT_NAME) $(GO_PATH)/bin/$(PROJECT_NAME)

docker: clean ## Build the Docker Image
	docker build -t $(NAMESPACE)/$(PROJECT_NAME):latest .

publish: docker ## Build and publish the Docker Image
	docker push $(NAMESPACE)/$(PROJECT_NAME):latest

lint: ## Lint the source code
	@cd ; go get golang.org/x/lint/golint
	@cd ; go get golang.org/x/tools/cmd/goimports
	go get -d ./...
	gofmt -s -w .
	go vet ./...
	$(GO_PATH)/bin/golint -set_exit_status=1 ./...
	$(GO_PATH)/bin/goimports -w .

test: ## Run the unit test and generate coverage
	@mkdir -p var/
	@go test -race -cover -coverprofile  var/coverage.txt ./...
	@go tool cover -func var/coverage.txt | awk '/^total/{print $$1 " " $$3}'

docs: ## Start a godoc server
	@cd ; go get golang.org/x/tools/cmd/godoc
	@echo "Docs here: http://localhost:3232/pkg/${MODULE}"
	@godoc -http=localhost:3232 -index -index_interval 2s -play

clean: ## Remove all git ignored file
	git clean -Xdf --exclude="!/.env"

githooks: ## Install the git hooks
	cp ops/hooks/* .git/hooks/
	chmod +x .git/hooks/*

dev: ## Run project in dev mode
	clear
	@go run main.go
