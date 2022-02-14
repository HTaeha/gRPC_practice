package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/HTaeha/gRPC_practice/ch3/client/ecommerce/order_management"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order, err := c.GetOrder(ctx, &pb.OrderID{Value: "first"})
	if err != nil {
		log.Fatalf("could not get order: %v", err)
	}
	log.Printf("GetOrder: %s", order)
}
