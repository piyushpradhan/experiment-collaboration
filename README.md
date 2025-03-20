# Collaboration Project

## Overview

The Collaboration Project is a real-time web application that allows users to connect and communicate through a WebSocket server. It provides a simple API for user management and supports CORS for cross-origin requests. This project is built using Go and follows a microservices architecture.

## Features

- Real-time communication using WebSockets
- User management API
- CORS support for specified origins
- In-memory storage for user data
- Modular architecture for easy maintenance and scalability

## Technologies Used

- Go (version 1.23.0)
- Gorilla WebSocket
- Go Kit for microservices
- UUID for unique identifiers

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.23.0 or higher)
- Git for version control

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/piyushpradhan/experiment-collaboration.git
   cd collaboration
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

### Running the Application

To run the application, use the following command:

```bash
go run main.go
```

The server will start on port 5000 by default. You can change the port by modifying the `listenAddr` flag.

### API Endpoints

- **GET /user**: Retrieve user information by ID.
- **DELETE /user/id**: Delete a user by ID.

### WebSocket Connection

To connect to the WebSocket server, use the following URL:

```
ws://localhost:5000/
```

## Directory Structure

```
.
├── api
│   ├── decode.go
│   ├── endpoint.go
│   ├── server.go
│   ├── service.go
│   └── transport.go
├── cmd
│   ├── api
│   │   └── main.go
│   └── collaboration
│       └── main.go
├── services
│   ├── api
│   │   ├── middleware.go
│   │   └── transport.go
│   └── collaboration
│       ├── endpoint.go
│       ├── service
│       │   └── service.go
│       └── transport.go
├── storage
│   └── memory.go
├── types
│   ├── client.go
│   ├── message.go
│   └── user.go
└── util
    └── util.go
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their support and resources.
- Inspired by various open-source projects and microservices architecture principles.
VB
