package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	proto "github.com/devesh/go-grpc-microservice/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listner, tcpErr := net.Listen("tcp", ":9001")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)
	if err := srv.Serve(listner); err != nil {
		panic(err)
	}

}

// proto file ke function ko implement kar rahe hai.
func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("Recieving request from client: ", req.Somestring)
	fmt.Println("Hello from server !")
	return &proto.HelloResponse{}, errors.New("")
}
