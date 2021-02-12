GO_PATH := $(shell go env GOPATH 2> /dev/null)
MODULE := $(shell awk '/^module/ {print $$2}' go.mod)
NAMESPACE := $(shell awk -F "/" '/^module/ {print $$(NF-1)}' go.mod)
PROJECT_NAME := $(shell awk -F "/" '/^module/ {print $$(NF)}' go.mod)
PATH := $(GO_PATH)/bin:$(PATH)

help:
	@echo "Makefile targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' Makefile \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/    \1 :: \3/p' \
	| column -t -c 1  -s '::'

full: clean lint test build ## Clean, test, and build the application

build: ## Build the application
	go build -o var/${PROJECT_NAME} .

docker: clean ## Build the Docker Image
	docker build -t $(NAMESPACE)/$(PROJECT_NAME):latest .

publish: docker ## Build and publish the Docker Image
	docker push $(NAMESPACE)/$(PROJECT_NAME):latest

watch: ## Run and auto-recompile on file changes
	@cd ; go get github.com/codegangsta/gin
	clear
	gin --all --immediate --path .. --build . --bin var/gin run

lint: ## Check the code for errors
	@cd ; go get golang.org/x/lint/golint
	@cd ; go get golang.org/x/tools/cmd/goimports
	go get -d ./...
	go mod tidy
	gofmt -s -w .
	go vet ./...
	golint -set_exit_status=1 ./...
	goimports -w .

test: ## Run the tests
	@mkdir -p var/
	@go test -race -cover -coverprofile  var/coverage.txt ./...
	@go tool cover -func var/coverage.txt | awk '/^total/{print $$1 " " $$3}'

clean: ## Remove all files listed in .gitignore
	git clean -Xdf

post-lint:
	@git diff --exit-code --quiet || (echo "There should not be any changes after the lint runs" && git status && exit 1;)

pipeline: full post-lint
