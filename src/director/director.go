package director

import "context"

type Daemons interface {
	Name() string
	Start(ctx context.Context, cancelFunc context.CancelFunc) error
	Stop(ctx context.Context) error
}
