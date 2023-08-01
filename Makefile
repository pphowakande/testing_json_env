.PHONY: test
test: 
	go test ./... -v -short -cover

build:
	go build