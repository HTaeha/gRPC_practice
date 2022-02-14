package ecommerce

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/HTaeha/gRPC_practice/ch3/server/ecommerce/order_management"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	orderMap map[string]*pb.Order
	pb.UnimplementedOrderManagementServer
}

func (s *Server) Init() {
	if s.orderMap == nil {
		s.orderMap = make(map[string]*pb.Order)
	}
	items := []string{"wejfo"}
	s.orderMap["first"] = &pb.Order{
		Id:          "0",
		Items:       items,
		Description: "description",
		Price:       3.14,
		Destination: "Destination",
	}
	s.orderMap["second"] = &pb.Order{
		Id:          "1",
		Items:       items,
		Description: "description",
		Price:       3.14,
		Destination: "Destination",
	}
}
func (s *Server) GetOrder(ctx context.Context, in *pb.OrderID) (*pb.Order, error) {
	// log.Printf("server: %s, %s", s.orderMap, in)
	value, exists := s.orderMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exist.", in.Value)
}

func (s *Server) SearchOrder(orderID *pb.OrderID, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(
				itemStr, searchQuery.Value
			){
				err := stream.Send(*order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}