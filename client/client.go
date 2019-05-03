package main

import (
	"context"
	"io"
	"log"

	pb "github.com/emilaasa/sse-demo/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}

	client := pb.NewSimpleProtoClient(conn)
	ctx := context.Background()

	stream, err := client.GetMessages(ctx,
		&pb.SimpleRequest{
			Name: "Tobias",
		})
	if err != nil {
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMessages(_) = _, %v", client, err)
		}
		log.Println(msg)
	}

}
