package ctxzap_test

import (
	"context"

	pb_testproto "github.com/grpc-ecosystem/go-grpc-middleware/grpctesting/testproto"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/interceptors/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/providers/zap/ctxzap"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

// Simple unary handler that adds custom fields to the requests's context. These will be used for all log statements.
func ExampleExtract_unary() {
	_ = func(ctx context.Context, ping *pb_testproto.PingRequest) (*pb_testproto.PingResponse, error) {
		// Add fields the ctxtags of the request which will be added to all extracted loggers.
		ctxtags.Extract(ctx).Set("custom_tags.string", "something").Set("custom_tags.int", 1337)

		// Extract a single request-scoped zap.Logger and log messages.
		l := ctxzap.Extract(ctx)
		l.Info("some ping")
		l.Info("another ping")
		return &pb_testproto.PingResponse{Value: ping.Value}, nil
	}
}
