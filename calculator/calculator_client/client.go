package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dueruen/gRPC-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Staring to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Staring to do a PrimeDeomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 121234567890,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDeomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
