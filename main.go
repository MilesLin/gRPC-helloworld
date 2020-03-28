package main

import (
	"google.golang.org/grpc"
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
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(s, &service.MessageService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
