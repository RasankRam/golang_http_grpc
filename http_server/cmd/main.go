package main

import (
	pb "example.com/task_platform_proto/gen_go"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
	"todo-list/internal/app/mw"
	"todo-list/internal/app/routes"
	"todo-list/internal/db"
	"todo-list/internal/logger"
	"todo-list/internal/utils"
	"todo-list/internal/validate"
)

func NewGRPCClient() (pb.TodoServiceClient, *grpc.ClientConn, error) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("grpc_server:5001", grpc.WithInsecure(), grpc.WithTimeout(60*time.Second))
	if err != nil {
		return nil, nil, err
	}

	// Create the client for the TodoService
	client := pb.NewTodoServiceClient(conn)
	return client, conn, nil
}

func main() {
	fmt.Println("Server is starting...")
	grpcClient, grpcConn, err := NewGRPCClient()
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	utils.LoadEnv(".env")
	validate.InitValidator()

	iLogger, file := logger.SetupLogger()
	defer file.Close()

	conn := db.OpenSqlxViaPgxConnPool()
	defer conn.Close()

	router := http.NewServeMux()

	routes.SetupTodoRoutes(router, conn, iLogger)
	routes.SetupAuthRoutes(router, conn, iLogger)

	wrappedMux := mw.SetGrpcClientMiddleware(router, grpcClient)

	http.ListenAndServe(":8080", wrappedMux)
}
