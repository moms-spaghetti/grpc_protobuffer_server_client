package main

import (
	"context"
	"log"
	"net"

	pb "github.com/moms-spaghetti/simple_grpc/proto"
	"google.golang.org/grpc"
)

const address = ":50051"

type server struct {
	pb.UnimplementedGreetingServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}