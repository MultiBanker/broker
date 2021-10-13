package middleware

import (
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	zerolog "github.com/philip-bui/grpc-zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func Options() []grpc.ServerOption {
	log := logger()
	return []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Second,
		}),
		zerolog.UnaryInterceptorWithLogger(&log),
		grpc.ChainUnaryInterceptor(
			grpcrecovery.UnaryServerInterceptor(),
		),
	}
}
