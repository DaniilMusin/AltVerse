package network

import (
    "net"

    "altverse/gameplay-engine/internal/game"
    pb "altverse/gameplay-engine/proto"
    "google.golang.org/grpc"
    "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
    "log/slog"
)

type Server struct {
    pb.UnimplementedGameplayServer
    engine *game.Engine
    log    *slog.Logger
}

func NewGRPCServer(e *game.Engine, log *slog.Logger) *Server {
    return &Server{engine: e, log: log}
}

func (s *Server) Start(addr string) {
    lis, err := net.Listen("tcp", addr)
    if err != nil { panic(err) }
    grpcServer := grpc.NewServer(
        grpc.StatsHandler(otelgrpc.NewServerHandler()),
    )
    pb.RegisterGameplayServer(grpcServer, s)
    if err := grpcServer.Serve(lis); err != nil { panic(err) }
}

func (s *Server) Stop() { /* … */ }
