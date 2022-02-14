package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/HTaeha/gRPC_practice/ch3/server/ecommerce"
	pb "github.com/HTaeha/gRPC_practice/ch3/server/ecommerce/order_management"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &ecommerce.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
