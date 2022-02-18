package main

import (
	"context"
	"flag"
	"io"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// // GetOrder
	// order, err := c.GetOrder(ctx, &pb.OrderID{Value: "first"})
	// if err != nil {
	// 	log.Fatalf("could not get order: %v", err)
	// }
	// log.Printf("GetOrder: %s", order)

	// // SearchOrders
	// searchStream, err := c.SearchOrders(ctx, &pb.OrderID{Value: "item2"})
	// if err != nil {
	// 	log.Fatalf("could not search orders: %v", err)
	// }
	// for {
	// 	searchOrder, err := searchStream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	log.Println("search result : ", searchOrder)
	// }

	// // UpdateOrders
	// updateStream, err := c.UpdateOrders(ctx)
	// if err != nil {
	// 	log.Fatalf("%v.UpdateOrders(_) = _, %v", c, err)
	// }

	// // Updating order 1
	// updOrder1 := &pb.Order{
	// 	Id:          "0",
	// 	Items:       []string{"a", "b", "c"},
	// 	Description: "description",
	// 	Price:       3.14,
	// 	Destination: "Destination",
	// }
	// if err := updateStream.Send(updOrder1); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	// }

	// // Updating order 2
	// updOrder2 := &pb.Order{
	// 	Id:          "1",
	// 	Items:       []string{"a2", "b2", "c2"},
	// 	Description: "description2",
	// 	Price:       123,
	// 	Destination: "Destination2",
	// }
	// if err := updateStream.Send(updOrder2); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	// }

	// // Updating order 3
	// updOrder3 := &pb.Order{
	// 	Id:          "200",
	// 	Items:       []string{"a3", "b3", "c3"},
	// 	Description: "description3",
	// 	Price:       1234,
	// 	Destination: "Destination3",
	// }
	// if err := updateStream.Send(updOrder3); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	// }

	// updateRes, err := updateStream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	// }
	// log.Printf("Update Orders Res : %s", updateRes)

	// ProcessOrders
	processOrdersStream, err := c.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", c, err)
	}
	if err := processOrdersStream.Send(&pb.OrderID{Value: "first"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "first", err)
	}
	if err := processOrdersStream.Send(&pb.OrderID{Value: "second"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "second", err)
	}

	channel := make(chan *pb.CombinedShipment)
	go asyncClientBidirectionalRPC(processOrdersStream, channel)
	time.Sleep(time.Millisecond * 1000)

	if err := processOrdersStream.Send(&pb.OrderID{Value: "second"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "second", err)
	}

	if err := processOrdersStream.CloseSend(); err != nil {
		log.Fatal(err)
	}
	// channel <- struct{}{}
	log.Printf("channel : %v", <-channel)
	log.Printf("channel : %v", <-channel)
}

func asyncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan *pb.CombinedShipment) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Printf("Combined shipment : %v", combinedShipment.OrdersList)
		c <- combinedShipment
	}
	// close(c)
	// c <- struct{}{}
}
