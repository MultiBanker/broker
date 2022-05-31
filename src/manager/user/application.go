package user

import "context"

type ApplicationManager interface {
	Create(ctx context.Context)
}
