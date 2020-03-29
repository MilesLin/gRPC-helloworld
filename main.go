package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"trygrpc/pb"
	"trygrpc/service"
)

const (
	port = ":50051"
)
func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")

	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(s, &service.MessageService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
