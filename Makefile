
build:
	@go build -o bin/goGui

run: build
	@./bin/goGui

test:
	@go test -v ./...