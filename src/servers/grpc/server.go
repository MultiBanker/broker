package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
)

type grpcServer struct {
	Address   string
	IsTesting bool

	manager manager.Abstractor
	server  *grpc.Server
	handler func(server *grpc.Server, abstract manager.Abstractor)
}

func (g *grpcServer) Name() string {
	return "GRPC"
}

func (g *grpcServer) Start(ctx context.Context, cancel context.CancelFunc) error {
	listener, err := net.Listen("tcp", g.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Second,
		}),
	}
	g.server = grpc.NewServer(opts...)

	g.handler(g.server, g.manager)

	log.Printf("[INFO] serving GRPC on \"%s\"", g.Address)
	return g.server.Serve(listener)
}

func (g *grpcServer) Stop(ctx context.Context) error {
	<-ctx.Done()
	g.server.GracefulStop()
	return nil
}

func NewGRPC(config *config.Config, man manager.Abstractor, handler func(server *grpc.Server, abstract manager.Abstractor)) *grpcServer {
	return &grpcServer{
		Address: config.GRPC.ListenAddr,
		manager: man,
		handler: handler,
	}
}

