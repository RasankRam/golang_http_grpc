package main

import (
	"context"
	pb "example.com/task_platform_proto/gen_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

// ProcessTodo возвращает сумму с вычетом налога, также проверяет, что имя задачи 6 и цена больше нуля
// It also validates that the name is at least 6 characters and price is at least 100.00
func (s *TodoServer) ProcessTodo(_ context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	discountedPrice := req.GetPrice() * 0.85

	roundedPrice := math.Round(discountedPrice*100) / 100

	res := &pb.TodoResponse{
		DiscountedPrice: roundedPrice,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, &TodoServer{})

	log.Println("gRPC server is running on port 5001...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
