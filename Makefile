.PHONY: build clean proto test

# Build the client application
build:
	@if [ ! -f proto/device.pb.go ] || [ ! -f proto/device_grpc.pb.go ]; then \
		echo "Generated protobuf files not found, running proto generation..."; \
		$(MAKE) proto; \
	fi
	go build -o starlink-client main.go

# Generate Go code from proto files
proto:
	@if ! command -v protoc >/dev/null 2>&1; then \
		echo "Error: protoc not found. Please install protobuf-compiler or use pre-generated files."; \
		echo "On Ubuntu/Debian: sudo apt install protobuf-compiler"; \
		echo "On macOS: brew install protobuf"; \
		exit 1; \
	fi
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
	@echo ""
	@echo "Note: If you plan to regenerate protobuf files, you also need protoc:"
	@echo "  Ubuntu/Debian: sudo apt install protobuf-compiler"
	@echo "  macOS: brew install protobuf"
	@echo "  The project includes pre-generated .pb.go files for convenience."

# Test the application (will fail without real device)
test: build
	go test -v
	./starlink-client -timeout 2s || echo "Test completed - expected connection failure without real device"

# Run unit tests only
unit-test:
	go test -v

# Run with default settings
run: build
	./starlink-client

# Show help
help:
	@echo "Available targets:"
	@echo "  build     - Build the client application"
	@echo "  proto     - Generate Go code from proto files"
	@echo "  clean     - Clean build artifacts"
	@echo "  deps      - Install dependencies"
	@echo "  test      - Run all tests (unit + integration)"
	@echo "  unit-test - Run unit tests only"
	@echo "  run       - Run with default settings"
	@echo "  help      - Show this help message"