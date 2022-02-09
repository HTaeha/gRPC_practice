package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/HTaeha/gRPC_practice/ch2/client/ecommerce/product_info"

	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to Product")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	description := "Description test"
	price := float32(25640.5)
	r, err := c.AddProduct(ctx, &pb.Product{Name: *name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("could not add product: %v", err)
	}
	log.Printf("AddProduct: %s", r.String())

	r2, err2 := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err2 != nil {
		log.Fatalf("could not get product: %v", err2)
	}
	log.Printf("GetProduct: %s", r2.String())
}
