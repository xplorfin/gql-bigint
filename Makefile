SERVICE_NAME=gql-bigint
# name of output binary
BINARY_NAME=server

# latest git commit hash
LATEST_COMMIT_HASH=$(shell git rev-parse HEAD)

# go commands and variables
GO=go
GOB=$(GO) build
GOM=$(GO) mod

# git commands
GIT=git

# environment variables related to
# cross-compilation.
GOOS_MACOS=darwin
GOOS_LINUX=linux
GOARCH=amd64

# currently installed/running Go version (full and minor)
GOVERSION=$(shell go version | grep -Eo '([1-9]\.[0-9]{1,3}\.*[0-9]{0,3})')
MINORVER=$(shell echo $(GOVERSION) | awk '{ split($$0, array, ".") } { print array[2] }')

GOPATH=$(shell go env gopath)
GOPATH_BIN=$(GOPATH)/bin

# Color code definitions
# Note: everything is bold.
GREEN=\033[1;38;5;70m
BLUE=\033[1;38;5;27m
LIGHT_BLUE=\033[1;38;5;32m
MAGENTA=\033[1;38;5;128m
RESET_COLOR=\033[0m

COLORECHO = $(1)$(2)$(RESET_COLOR)

check-binary = $(shell which $(1) > /dev/null 2>&1; echo "$$?")
check-env = $(shell [ -n $$$(1) ]; echo "$$?")
check-minorver = $(shell [ $(MINORVER) -gt 15 ]; echo "$$?")

default: help

all: mod-tidy mod-vendor lint test

git-hooks: ## setup the repository (enables git hooks and downloads dependencies)
	git clone --depth=1 -q git@github.com:xplorfin/githooks-go .github/hooks
	@rm -rf .github/hooks/.git
	git config core.hooksPath .github/hooks --replace-all

update-hooks:
	rm -rf .github/hooks
	@make git-hooks

## go mod recipes

mod-vendor:
	$(GOM) vendor

mod-clean: ## Remove all the Go mod cache
	@go clean -modcache

mod-tidy:
	$(GOM) tidy

clean-mods: mod-clean

godocs: ## Run a godoc server
	@echo "godoc server running on http://localhost:9000"
	@godoc -http=":9000"

## testing/linting recipes
lint: ## Run the golangci-lint application (install if not found) & fix issues if possible
	@golangci-lint run --fix

ci-lint: lint

lint-hooks: git-hooks lint

test: ## run tests without coverage reporting
	@cd ./gql-test && go test ./...

ci-test: install-goacc # run a test with coverage
	@cd ./gql-test && go-acc -o profile.cov ./...

coverage: ## Get the test coverage from go-coverage
	@go test -coverprofile=coverage.out ./gql-test && go tool cover -func=coverage.out

bench:  ## Run all benchmarks in the Go application
	@go test -bench=. -benchmem

ci-all: ci-lint ci-test

## tool installer recipes
install-tools: install-goacc

install-goacc: TOOL_BINARY=go-acc
install-goacc: TOOL_REPO=github.com/ory/go-acc
install-goacc: install-gotool

install-gotool:
ifeq ($(call check-binary,$(TOOL_BINARY)), 1)
ifeq ($(call check-minorver), 0)
	@## go install is recommended/not deprecated for go 1.16 and above
	go install $(TOOL_REPO)@latest
else
	go get -u $(TOOL_REPO)
endif
endif

# pre-commit hook
pre-commit: lint-hooks

# build recipes
build-macos: BUILD_OS=$(GOOS_MACOS)
build-macos: build

build-linux: BUILD_OS=$(GOOS_LINUX)
build-linux: build

build:
	env GOOS=$(BUILD_OS) GOARCH=$(GOARCH) \
	$(GOB) -mod vendor -o $(BINARY_NAME)

help: ## This help dialog.
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "%-30s %s\n" $$help_command $$help_info ; \
	done
