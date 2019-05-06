package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	pb "github.com/emilaasa/sse-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type service struct {
	eventID int
}

func (s *service) GetMessages(req *pb.SimpleRequest, stream pb.SimpleProto_GetMessagesServer) error {
	headers, _ := metadata.FromIncomingContext(stream.Context())
	fmt.Println(headers)
	for {
		time.Sleep(2000 * time.Millisecond)
		s.eventID++
		if err := stream.Send(
			&pb.SimpleResponse{
				EventID: strconv.Itoa(s.eventID),
				Payload: "some message",
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

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported!", http.StatusInternalServerError)
		return
	}
	reqdump, err := httputil.DumpRequest(r, false)
	if err != nil {
		log.Fatalf("err: %v failed to dump request, paniiic!", err)
	}

	fmt.Println(string(reqdump))
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}

	client := pb.NewSimpleProtoClient(conn)

	var ctx context.Context
	id := r.Header.Get("Last-Event-Id")
	if id != "" {
		header := metadata.New(map[string]string{"Last-Event-Id": id})
		ctx = metadata.NewOutgoingContext(context.Background(), header)
	} else {
		ctx = context.Background()
	}

	stream, _ := client.GetMessages(ctx,
		&pb.SimpleRequest{})

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Fprint(w, "\n\n")
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMessages(_) = _, %v", client, err)
		}
		fmt.Fprintf(w, "id: %s\ndata: %s\n\n", msg.EventID, msg.Payload)
		f.Flush()
	}
}

func main() {
	go func() {
		grpcServer()
	}()
	fs := http.FileServer(http.Dir("./debug"))
	http.Handle("/debug/", http.StripPrefix("/debug/", fs))
	http.HandleFunc("/", proxyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
