# Experiment Collaboration Server

A real-time collaboration server built with Go that enables concurrent editing and collaboration features through WebSocket connections.

## Features

- Real-time collaboration through WebSocket connections
- In-memory storage for managing collaboration sessions
- RESTful API endpoints for session management
- Concurrent editing support
- Lightweight and efficient design

## Prerequisites

- Go 1.23.0 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/experiment-collaboration.git
cd experiment-collaboration
```

2. Install dependencies:
```bash
go mod download
```

## Project Structure

```
.
├── api/        # API handlers and server implementation
├── cmd/        # Command-line tools and entry points
├── services/   # Business logic and service layer
├── storage/    # Storage implementations (memory storage)
├── types/      # Type definitions and interfaces
├── util/       # Utility functions and helpers
├── websocket/  # WebSocket handling and real-time communication
├── main.go     # Application entry point
└── .air.toml   # Air configuration for live reloading
```

## Running the Server

1. Start the server with default settings:
```bash
go run main.go
```

2. Start the server with a custom port:
```bash
go run main.go -listenaddr :8080
```

The server will start and listen on the specified port (default: 5000).

## Development

For development with live reloading, you can use Air:

```bash
air
```

This will automatically rebuild and restart the server when files change.

## Dependencies

- [github.com/google/uuid](https://github.com/google/uuid) - For generating unique identifiers
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket) - For WebSocket implementation
