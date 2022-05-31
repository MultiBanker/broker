package repository

import "context"

type Sequencer interface {
	NextSequenceValue(ctx context.Context, sequenceName string) (int, error)
	NewSequence(ctx context.Context, sequenceName string) error
}
