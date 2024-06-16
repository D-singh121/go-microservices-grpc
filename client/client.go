package main

import (
	"context"

	proto "github.com/devesh/go-grpc-microservice/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewExampleClient(conn)
	req := &proto.HelloRequest{Somestring: "Hello from client!"}
	client.ServerReply(context.TODO(), req)
}
