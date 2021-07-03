package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

// SayHello Is this fine
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Receieved Message Body from client: %s", message.Body)
	return &Message{Body: "Hello from server"}, nil
}
