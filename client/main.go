package main

import (
	"context"
	"fmt"

	pb "github.com/mridulganga/gRPC-golang-sample"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	c := pb.NewMathClient(conn)

	out, err := c.Add(context.Background(), &pb.NumInput{
		N1: 3.14, N2: 2.14,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	out, err = c.Substract(context.Background(), &pb.NumInput{
		N1: 3.14, N2: 2.14,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

}
