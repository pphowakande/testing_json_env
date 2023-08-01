.PHONY: test
test: 
	gotestdox -short -count=1 ./...

.PHONY: lint
lint: 
	golangci-lint run