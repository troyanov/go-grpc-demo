package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	hellopb "github.com/troyanov/go-grpc-demo/proto/hello"
)

type server struct {
	hellopb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	return &hellopb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	hellopb.RegisterGreeterServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(listener))
}
