package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/moms-spaghetti/simple_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = ":50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Ani"})
	if err != nil {
		log.Fatalf("unable to greet: %v", err)
	}
	fmt.Println("greeting: " + r.GetMessage())
}