export GOLANGCI_LINT_VERSION := v1.36.0

.PHONY: test
test:
	@echo "<-- Running Tests -->"
	@go test -v ./... | grep -v "no test files"

.PHONY: githooks-init
githooks-init:
	cp .pre-commit .git/hooks/pre-commit

.PHONY: lint
lint: bin/golangci-lint
	@./bin/golangci-lint run ./...

bin/golangci-lint:
	@curl -fsSL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s $(GOLANGCI_LINT_VERSION)
