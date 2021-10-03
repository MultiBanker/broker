package repository

import "context"

func (r Repository) SequenceRepo() SequencesRepository {
	return r.Sequence
}

type SequencesRepository interface {
	NextSequenceValue(ctx context.Context, sequenceName string) (int, error)
	NewSequence(ctx context.Context, sequenceName string) error
}

