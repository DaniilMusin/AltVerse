package proto

import "google.golang.org/grpc"

type UnimplementedGameplayServer struct{}

type GameplayServer interface{}

func RegisterGameplayServer(s *grpc.Server, srv GameplayServer) {}

type SubscribeRequest struct{ PlayerId string }
type PlayerCommand struct{ PlayerId string; Payload string }
type CommandAck struct{ Id string; Accepted bool }
type GameEvent struct{ Topic string; Data []byte; Ts int64 }
