package main

import (
	"context"
	"fmt"
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
	server.Serve(lis)

	service := &service{}
	pb.RegisterSimpleProtoServer(server, service)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}

func grpcClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}

	client := pb.NewSimpleProtoClient(conn)
	ctx := context.Background()

	rsp, err := client.GetMessage(ctx, &pb.SimpleRequest{
		Name: "Tobias",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rsp= %+v\n", rsp)

}

func main() {
	go grpcServer()
	grpcClient()
}
