.PHONY: test
test: 
	go test ./... -v -short -cover

.PHONY: lint
lint: 
	golangci-lint run