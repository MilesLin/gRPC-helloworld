package service

import (
	"context"
	"trygrpc/pb"
)

type MessageService struct{
}

func (m MessageService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}

