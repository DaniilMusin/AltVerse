package network

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func UnaryLogger(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now()
		span := trace.SpanFromContext(ctx)
		resp, err = handler(ctx, req)
		log.Info("grpc", "method", info.FullMethod, "dur", time.Since(start), "err", err,
			"trace_id", span.SpanContext().TraceID())
		return
	}
}
