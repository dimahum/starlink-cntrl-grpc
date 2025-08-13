# Starlink gRPC Control Client

A Go client for connecting to Starlink devices via gRPC and reading status information.

## Features

- Connect to Starlink devices using gRPC
- Read device status including:
  - Device information (ID, hardware/software versions, country code)
  - Device state (uptime)
- Configurable connection parameters
- Support for both secure and insecure connections

## Prerequisites

- Go 1.19 or later
- Access to a Starlink device on the local network

## Building

### Local Build

```bash
go mod tidy
go build -o starlink-client main.go
```

### Cross-Platform Build

Using the Makefile:
```bash
make build-cross GOOS=linux GOARCH=amd64
make build-cross GOOS=darwin GOARCH=amd64
make build-cross GOOS=windows GOARCH=amd64
```

### Automated Builds

The project includes GitHub Actions workflows that automatically:
- Build and test on every push to main branch and pull requests
- Create cross-platform binaries (Linux, macOS, Windows for AMD64 and ARM64)
- Publish releases with binaries when git tags are created

To create a release:
```bash
git tag v1.0.0
git push origin v1.0.0
```

## Usage

Basic usage (connects to default Starlink IP):
```bash
./starlink-client
```

Custom address:
```bash
./starlink-client -addr 192.168.100.1:9200
```

With custom timeout:
```bash
./starlink-client -timeout 30s
```

For secure connections:
```bash
./starlink-client -insecure=false
```

## Command Line Options

- `-addr`: Starlink device address (default: "192.168.100.1:9200")
- `-insecure`: Use insecure connection (default: true)
- `-timeout`: Connection timeout (default: 10s)

## Default Starlink Network Configuration

Starlink devices typically:
- Use IP address `192.168.100.1` when in bypass mode
- Expose gRPC service on port `9200`
- May require insecure connections depending on firmware

## Protocol Buffers

The application uses Protocol Buffers to define the gRPC service interface. The proto file is located in `proto/device.proto` and includes definitions for:

- Device service with Handle method
- GetStatus request/response messages
- Device information and state structures

## License

MIT License - see LICENSE file for details.