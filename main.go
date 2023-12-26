package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection from client ", ws.RemoteAddr())
	s.connections[ws] = true

	s.readLoop(ws)
}

// readLoop continuously reads messages from the websocket connection and broadcasts them to all connected clients.
// It also sends a confirmation message back to the sender.
func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		msg := string(buf[:n])
		fmt.Println("Received message: ", msg)
		ws.Write([]byte("thanks, recieved your message"))
		s.broadcast([]byte(msg))
	}
}

func (s *Server) broadcast(b []byte) {
	for conn := range s.connections {
		go func(c *websocket.Conn) {
			_, err := c.Write(b)
			if err != nil {
				log.Println(err)
				return
			}
		}(conn)
	}
}

func main() {
	server := NewServer()
	http.Handle("/", websocket.Handler(server.handleWS))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
