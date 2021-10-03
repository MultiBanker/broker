package servers

import "context"

type Server interface {
	Name() string
	Start(ctx context.Context, cancel context.CancelFunc) error
	Stop(ctx context.Context) error
}
