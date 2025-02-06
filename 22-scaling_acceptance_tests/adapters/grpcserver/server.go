package grpcserver

import (
	"context"
	interactions "go-specs-greet/v2/domain/interactions"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func (g GreetServer) Introduce(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Introduce(request.Name)}, nil
}
