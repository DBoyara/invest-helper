PROGRAM_NAME = invest-bot

.PHONY: help clean dep test build build-docker

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

clean: ## Clean build directory.
	rm -f ./bin/${PROGRAM_NAME}
	rmdir ./bin

dep: ## Download the dependencies.
	go mod download

test: ## Run tests
	go test -v -cover ./...

go-lint: ## Run linter
	@docker run --rm -it -v `pwd`:/go/src/invest-bot -w /go/src/invest-bot golangci/golangci-lint golangci-lint run

lint: ## Run linter
	go vet ./...

build: dep ## Build executable.
	go build -v main.go

build-docker: dep ## Build docker.
	go build -o invest-bot main.go

migrate:
	@dbmate -u "postgres://user:pass@localhost:5433/invest-bot?sslmode=disable" --no-dump-schema up
