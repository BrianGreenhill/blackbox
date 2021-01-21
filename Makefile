.PHONY: test
test:
	@echo "<-- Running Tests -->"
	@go test -mod vendor $(shell go list -mod vendor ./...) | grep -v "no test files"
