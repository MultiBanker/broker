package repository

import "context"

func (r Repository) SequenceRepo() Sequencer {
	return r.Sequence
}

type Sequencer interface {
	NextSequenceValue(ctx context.Context, sequenceName string) (int, error)
	NewSequence(ctx context.Context, sequenceName string) error
}

