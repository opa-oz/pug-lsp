.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run unit tests
	@go test ./pkg/utils ./pkg/query ./pkg/html ./pkg/documents

.PHONY: test

build: ## Build GO binary
	@go build -o pug-lsp -ldflags="-s -w"

.PHONY: build

toc: ## Format README.md to add TOC
	markdown-toc -i README.md

.PHONY: toc
