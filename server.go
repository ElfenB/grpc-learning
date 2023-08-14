package main

import (
	"log"
	"net"

	"chat/chat"

	"google.golang.org/grpc"
)

func main() {
	PORT := ":8080"

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("Cannot create listener on port %s: %s", PORT, err)
	}

	serverRegistrar := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &serverRegistrar)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to launch server on port %s: %s", PORT, err)
	}
}
