package grpc

import (
	"context"
	"log"
	"net"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/grpc/middleware"
	"google.golang.org/grpc"
)

type grpcServer struct {
	Address   string
	IsTesting bool

	manager manager.Managers
	server  *grpc.Server
	handler func(server *grpc.Server, abstract manager.Managers)
}

func (g *grpcServer) Name() string {
	return "GRPC"
}

func (g *grpcServer) Start(_ context.Context, cancel context.CancelFunc) error {
	defer cancel()
	listener, err := net.Listen("tcp", g.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	g.server = grpc.NewServer(middleware.Options()...)

	g.handler(g.server, g.manager)

	return g.server.Serve(listener)
}

func (g *grpcServer) Stop(_ context.Context) error {
	g.server.GracefulStop()
	return nil
}

func NewGRPC(config *config.Config, man manager.Managers, handler func(server *grpc.Server, abstract manager.Managers)) *grpcServer {
	return &grpcServer{
		Address: config.GRPC.ListenAddr,
		manager: man,
		handler: handler,
	}
}
