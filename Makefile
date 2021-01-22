.PHONY: test
test:
	@echo "<-- Running Tests -->"
	@go test -v ./... | grep -v "no test files"
