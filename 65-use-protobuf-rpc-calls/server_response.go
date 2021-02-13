package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Masterminds/go-in/practice/chapter10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloREsponse, error) {
	msg := "Hello " + in.Name + "!"
	return &pb.HelloResponse{Message: msg}, nil
}

func main() {
	l, err := net.Listen("tcp", ":5555")
	if err != nil {

	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
