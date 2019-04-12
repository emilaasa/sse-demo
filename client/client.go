package main

import (
	"context"
	"fmt"
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

	rsp, err := client.GetMessage(ctx, &pb.SimpleRequest{
		Name: "Tobias",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rsp = %+v\n", rsp)

}
