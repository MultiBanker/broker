package loan

import (
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
)

func DtoToModelLoanProgram(req dto.LoanProgramRequest) models.LoanProgram {
	return models.LoanProgram{
		Name:        req.Name,
		Note:        req.Note,
		Code:        req.Code,
		IsEnabled:   req.IsEnabled,
		MaxAmount:   req.MaxAmount,
		MinAmount:   req.MinAmount,
		PartnerCode: req.PartnerCode,
		Rate:        req.Rate,
		Term:        req.Term,
		Type:        req.Type,
	}
}
