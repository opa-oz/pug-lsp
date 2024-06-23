PACKAGE_NAME          := github.com/opa-oz/pug-lsp
GOLANG_CROSS_VERSION  ?= v1.21.5

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

.PHONY: release-dry-run
release-dry-run:
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v `pwd`/sysroot:/sysroot \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--clean --skip=validate --skip=publish
