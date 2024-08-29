
APP_NAME = "monkey-compiler"

.PHONY: run format build

run: build
	@./bin/$(APP_NAME)

format:
	@go fmt ./...

build: format
	@echo "Building $(APP_NAME)"
	@go build -o bin/$(APP_NAME) cmd/main.go

test:
	# @go test ./... -v
	# test the internal folders
	@go test ./internal/... -v

playground:
	@go build -o bin/$(APP_NAME) cmd/main.go && ./bin/$(APP_NAME)

compare:
	@go build -o bin/fibonacci ./internal/benchmark && ./bin/fibonacci -engine=eval -fib=20 && ./bin/fibonacci -engine=vm -fib=20
