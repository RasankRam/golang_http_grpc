package mw

import (
	"context"
	pb "example.com/task_platform_proto/gen_go"
	"net/http"
	"todo-list/internal/utils"
)

func SetGrpcClientMiddleware(next http.Handler, grpcClient pb.TodoServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the gRPC client to the request context
		ctx := context.WithValue(r.Context(), utils.GrpcClientKey, grpcClient)

		// Pass the request with the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
