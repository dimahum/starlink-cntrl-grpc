.PHONY: build clean proto test

# Build the client application
build: proto
	go build -o starlink-client main.go

# Generate Go code from proto files
proto:
	export PATH=$$PATH:$$(go env GOPATH)/bin && \
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       proto/device.proto

# Clean build artifacts
clean:
	rm -f starlink-client
	rm -f proto/*.pb.go

# Install dependencies
deps:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Test the application (will fail without real device)
test: build
	./starlink-client -timeout 2s || echo "Test completed - expected connection failure without real device"

# Run with default settings
run: build
	./starlink-client

# Show help
help:
	@echo "Available targets:"
	@echo "  build  - Build the client application"
	@echo "  proto  - Generate Go code from proto files"
	@echo "  clean  - Clean build artifacts"
	@echo "  deps   - Install dependencies"
	@echo "  test   - Test the application"
	@echo "  run    - Run with default settings"
	@echo "  help   - Show this help message"