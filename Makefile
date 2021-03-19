GO := go
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v generated.go | grep -v _gen.go)

gql-generate: ## Regenerate GraphQL code
	@echo "+ $@"
	@go run github.com/99designs/gqlgen generate
.PHONY: gql-generate

db-generate: ## Generate SQL code with sqlc
	@echo "+ $@"
	@sqlc generate
.PHONY: db-generate