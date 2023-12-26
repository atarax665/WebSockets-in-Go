# WebSocket PubSub Server

This is a simple WebSocket PubSub (Publish-Subscribe) server implemented in Go, using the `golang.org/x/net/websocket` package.

## Overview

The server allows clients to connect via WebSocket and perform the following actions:

- **Publish**: Send a message to a specific topic.
- **Subscribe**: Subscribe to receive messages from a specific topic.
- **Unsubscribe**: Unsubscribe from a specific topic.

Clients are identified by a unique ID generated upon connection.

## Getting Started

### Prerequisites

- Go installed on your machine
- [Satori Go UUID](https://github.com/satori/go.uuid) package

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repository.git
   ```

2. Install dependencies:

   ```bash
   go get -u golang.org/x/net/websocket
   go get -u github.com/satori/go.uuid
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

The server will start running at [http://localhost:3000/ws](http://localhost:3000/ws).

## Usage

1. Open the provided `static/index.html` file in a web browser.

2. Connect to the WebSocket server.

3. Use the form on the page to perform actions:

   - **Action**: Enter one of the following actions: `publish`, `subscribe`, or `unsubscribe`.
   - **Topic**: Enter the topic for the action.
   - **Message**: Enter the message for the `publish` action.

4. Click the "Send" button to perform the action.

## Actions

### Publish

- **Action:** `publish`
- **Topic:** Specify the topic for the message.
- **Message:** Specify the message to be published to the topic.

### Subscribe

- **Action:** `subscribe`
- **Topic:** Specify the topic to subscribe to.

### Unsubscribe

- **Action:** `unsubscribe`
- **Topic:** Specify the topic to unsubscribe from.

## Notes

- The server logs actions, subscription status, and disconnections to the console.
- Clients are identified by a unique ID generated on connection.

Feel free to extend and modify this server for your specific use case or integrate it into your application.
