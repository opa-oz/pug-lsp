.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run unit tests
	@go test ./pkg/utils ./pkg/query ./pkg/html ./pkg/documents ./pkg/lsp ./pkg/completion

.PHONY: test


cover: ## Run tests and collect coverage
	@go test -cover ./pkg/utils ./pkg/query ./pkg/html ./pkg/documents ./pkg/lsp ./pkg/completion

.PHONY: cover

build: ## Build GO binary
	@go build -o pug-lsp -ldflags="-s -w"

.PHONY: build

toc: ## Format README.md to add TOC
	markdown-toc -i README.md

.PHONY: toc
