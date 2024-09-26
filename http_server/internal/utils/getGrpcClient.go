package utils

import (
	"context"
	pb "example.com/task_platform_proto/gen_go"
)

type contextKey string

const GrpcClientKey contextKey = "grpcClient"

func GetGrpcClientFromContext(ctx context.Context) (pb.TodoServiceClient, bool) {
	grpcClient, ok := ctx.Value(GrpcClientKey).(pb.TodoServiceClient)
	return grpcClient, ok
}
