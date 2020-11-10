##@ Go Command

run:  ## Go run
	go run ./cmd/example.go

test:   ## Run go test
	go test -v --count=1 ./...

.PHONY: run test


##@ Lint

install-lint:  ## Install golangci-lint binary to ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.32.2

lint-testing:  ## Run lint
	./bin/golangci-lint run ./...

.PHONY: install-lint lint-testing

##@ Help

.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
