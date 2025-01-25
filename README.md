# Proglog - Distributed Log Service

A distributed commit log service implementation in Go, featuring both gRPC and HTTP APIs for log operations.

## Overview

Proglog is a distributed logging service that provides:
- Persistent append-only log storage
- gRPC and HTTP interfaces
- Streaming support for real-time log production and consumption
- Memory-mapped file operations for efficient storage

## Features

- **Commit Log**: Append-only log implementation with persistent storage
- **Multiple Interfaces**:
  - gRPC API for high-performance client-server communication
  - HTTP API for RESTful interactions
- **Streaming Support**: Real-time log production and consumption
- **Efficient Storage**: Memory-mapped file operations for optimal performance

## Getting Started

### Prerequisites

- Go 1.23.1 or later
- Protocol Buffers compiler (protoc)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/natnael-alemayehu/proglog.git
cd proglog
```

2. Install dependencies:
```bash
go mod download
```

### Running the Server

Start the server using:
```bash
go run cmd/server/main.go
```

The server will start on port 8000 by default.

## API Reference

### gRPC API

The service provides the following gRPC endpoints:

- `Produce`: Append a record to the log
- `Consume`: Read a record from the log
- `ProduceStream`: Stream records to the log
- `ConsumeStream`: Stream records from the log

### HTTP API

The service also provides HTTP endpoints for RESTful interactions.

## Project Structure

```
.
├── api/v1/            # Protocol Buffers definitions and generated code
├── cmd/
│   └── server/        # Server entry point
├── internal/
│   ├── log/          # Core log implementation
│   └── server/       # Server implementation (gRPC and HTTP)
```

## Dependencies

- `google.golang.org/grpc` - gRPC framework
- `github.com/gorilla/mux` - HTTP router
- `github.com/tysonmote/gommap` - Memory-mapped file operations
- `google.golang.org/protobuf` - Protocol Buffers support

## License

This project is licensed under the MIT License - see the LICENSE file for details. 