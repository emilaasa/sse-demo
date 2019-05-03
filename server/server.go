package main

import (
	"fmt"
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

	for {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "data: Message: %s\n\n", "LOL")
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
