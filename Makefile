-include .env
export

GO_PATH := $(shell go env GOPATH)
PROJECT_NAME := $(shell awk -F "/" '/^module/ {print $$(NF)}' go.mod)
NAMESPACE := $(shell awk -F "/" '/^module/ {print $$(NF-1)}' go.mod)
MODULE := $(shell awk '/^module/ {print $$2}' go.mod)

release: githooks clean lint test build

build:
	go build -o $(CURDIR)/var/$(PROJECT_NAME)
	@ln -sf $(CURDIR)/var/$(PROJECT_NAME) $(GO_PATH)/bin/$(PROJECT_NAME)

docker: clean
	docker build -t $(NAMESPACE)/$(PROJECT_NAME):latest .

publish: docker
	docker push $(NAMESPACE)/$(PROJECT_NAME):latest

lint:
	@cd ; go get golang.org/x/lint/golint
	@cd ; go get golang.org/x/tools/cmd/goimports
	go get -d ./...
	gofmt -s -w .
	go vet ./...
	$(GO_PATH)/bin/golint -set_exit_status=1 ./...
	$(GO_PATH)/bin/goimports -w .

test:
	@mkdir -p var/
	@go test -race -cover -coverprofile  var/coverage.txt ./...
	@go tool cover -func var/coverage.txt | awk '/^total/{print $$1 " " $$3}'

docs:
	@cd ; go get golang.org/x/tools/cmd/godoc
	@echo "Docs here: http://localhost:3232/pkg/${MODULE}"
	@godoc -http=localhost:3232 -index -index_interval 2s -play

clean:
	git clean -Xdf --exclude="!/.env"

githooks:
	cp ops/hooks/* .git/hooks/
	chmod +x .git/hooks/*

dev:
	clear
	@go run main.go
