.PHONY: test
test:
	@echo "<-- Running Tests -->"
	@cd backend && \
	go test -mod vendor $(shell cd backend && go list -mod vendor ./...) | grep -v "no test files"
