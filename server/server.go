package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/emilaasa/sse-demo/proto"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) GetMessages(req *pb.SimpleRequest, stream pb.SimpleProto_GetMessagesServer) error {
	for {
		time.Sleep(2000 * time.Millisecond)
		if err := stream.Send(
			&pb.SimpleResponse{
				Name: "Emil",
			}); err != nil {
			return err
		}
	}
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

func handler(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

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
			fmt.Fprint(w, "\n\n")
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMessages(_) = _, %v", client, err)
		}
		fmt.Fprintf(w, "data: %s\n\n", msg)
		f.Flush()
	}
}

func main() {
	go func() {
		grpcServer()
	}()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Look at the grpc headers
// Can I do the event-id in message or out of message?
// Resume the stream from an id
// SAD paths:
// 1. Event ID is missing in the event queue
