package main

import (
	"context"
	"log"
	"net"

	pb "github.com/emilaasa/sse-demo/proto"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) GetMessage(ctx context.Context, r *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	return &pb.SimpleResponse{
		Name: "Emil",
	}, nil
}

func grpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service := &service{}
	pb.RegisterSimpleProtoServer(server, service)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
func main() {
	grpcServer()
}
