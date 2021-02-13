package main

import (
	"context"
	"fmt"
	pb "github.com/Masterminds/go-in-practice/chapter10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:55555"
	conn, err := grpc.Dial(address, grpc.WithInsercure())
	if err != nil {

	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := "Inigo Montoya"
	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {

	}
	fmt.Println(r.Message)
}
