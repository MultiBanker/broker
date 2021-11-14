package loan

import (
	"context"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

var _ Program = (*ProgramManager)(nil)

type ProgramManager struct {
	loanColl repository.LoanProgram
	seqColl  repository.Sequencer
}

func NewProgramManager(repo repository.Repositories) *ProgramManager {
	return &ProgramManager{
		loanColl: repo.LoanProgramRepo(),
		seqColl:  repo.SequenceRepo(),
	}
}

type Program interface {
	LoanProgram(ctx context.Context, code string) (models.LoanProgram, error)
	LoanPrograms(ctx context.Context, paging selector.Paging) ([]models.LoanProgram, int64, error)
	CreateLoanProgram(ctx context.Context, program models.LoanProgram) (string, error)
	UpdateLoanProgram(ctx context.Context, code string, program models.LoanProgram) error
}

func (p ProgramManager) LoanProgram(ctx context.Context, code string) (models.LoanProgram, error) {
	return p.loanColl.LoanProgram(ctx, code)
}

func (p ProgramManager) LoanPrograms(ctx context.Context, paging selector.Paging) ([]models.LoanProgram, int64, error) {
	return p.loanColl.LoanPrograms(ctx, paging)
}

func (p ProgramManager) CreateLoanProgram(ctx context.Context, program models.LoanProgram) (string, error) {
	idInt, err := p.seqColl.NextSequenceValue(ctx, models.LoanSequences)
	if err != nil {
		return "", err
	}
	program.ID = strconv.Itoa(idInt)
	return p.loanColl.CreateLoanProgram(ctx, program)
}

func (p ProgramManager) UpdateLoanProgram(ctx context.Context, code string, program models.LoanProgram) error {
	return p.loanColl.UpdateLoanProgram(ctx, code, program)
}
