.PHONY: build run test clean dev install

# Build the server binary
build:
	go build -o bin/taskwarrior-api ./cmd/server

# Run the server
run:
	go run ./cmd/server

# Run with hot reload (requires air: go install github.com/air-verse/air@latest)
dev:
	air

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Install dependencies
install:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the server binary"
	@echo "  run            - Run the server"
	@echo "  dev            - Run with hot reload (requires air)"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  install        - Install/update dependencies"
	@echo "  clean          - Clean build artifacts"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code (requires golangci-lint)"

