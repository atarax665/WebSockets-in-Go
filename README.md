# Go WebSocket Server

I have implemented various nuances of web sockets using Golang, each implementation will be added in a new branch. Do check it out.

## 1. Simple websocket server

This is a simple WebSocket server written in Go using the `golang.org/x/net/websocket` package. The server allows multiple clients to connect and exchange messages over WebSocket.

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [WebSocket Communication](#websocket-communication)
- [License](#license)

## Overview

This WebSocket server is designed to handle multiple client connections concurrently. Each connected client can send messages to the server, and the server will broadcast those messages to all connected clients. The server also sends a confirmation message back to the sender for each received message.

## Getting Started

### Prerequisites

Before running the server, make sure you have the following installed:

- Go (version 1.11 or later)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/atarax665/WebSockets-in-Go.git
   cd WebSockets-in-Go
   ```

2. Run the server:

   ```bash
   go run main.go
   ```

By default, the server will run on port 8080. You can change the port in the `main` function of the `main.go` file.

## Usage

1. Run the server as described in the [Installation](#installation) section.

2. Open your WebSocket-supported client (e.g., a browser with WebSocket support or a WebSocket client).

3. Connect to the WebSocket server at `ws://localhost:8080` (or the configured port).

4. Send messages to the server, and observe the server broadcasting those messages to all connected clients.

5. The server will acknowledge each received message with a "thanks, received your message" response.

## WebSocket Communication

- **Connecting to the WebSocket server:**
  Connect to `ws://localhost:8080` using your WebSocket client.

- **Sending messages:**
  Send messages to the server, and the server will broadcast them to all connected clients.

- **Receiving messages:**
  Each connected client will receive broadcasted messages from other clients.

- **Acknowledgment:**
  The server sends a confirmation message back to the sender for each received message.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
