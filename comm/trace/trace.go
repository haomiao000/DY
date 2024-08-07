package trace

import (
	"context"

	"google.golang.org/grpc"
)

func ServerFilte(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	return handler(ctx, req)
}
