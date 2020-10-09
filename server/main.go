package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/mridulganga/gRPC-golang-sample"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.NumInput) (*pb.NumOutput, error) {
	fmt.Printf("Add %+v\n", in)
	return &pb.NumOutput{
		Ans: in.GetN1() + in.GetN2(),
	}, nil
}

func (s *server) Substract(ctx context.Context, in *pb.NumInput) (*pb.NumOutput, error) {
	fmt.Printf("Substract %+v\n", in)
	return &pb.NumOutput{
		Ans: in.GetN1() - in.GetN2(),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	serverStruct := server{}

	s := grpc.NewServer()

	pb.RegisterMathServer(s, &serverStruct)

	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
