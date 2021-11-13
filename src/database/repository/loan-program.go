package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) LoanProgramRepo() LoanProgram {
	return r.LoanProgram
}

type LoanProgram interface {
	LoanProgram(ctx context.Context, code string) (models.LoanProgram, error)
	LoanPrograms(ctx context.Context, paging selector.Paging) ([]models.LoanProgram, int64, error)
	CreateLoanProgram(ctx context.Context, program models.LoanProgram) (string, error)
	UpdateLoanProgram(ctx context.Context, code string, program models.LoanProgram) error
}
