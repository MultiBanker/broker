package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// @Router /admins/loan-programs/ [post]
// @Summary Создание кредитной программы
// @Description Создание кредитной программы
// @Tags programs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "токен авторизации"
// @Param program body dto.LoanProgramRequest true "кредитная программа"
// @Success 200 {object} models.Response
// @Failure 400 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
func (res AdminResource) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	var req dto.LoanProgramRequest

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err = req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.loanMan.CreateLoanProgram(ctx, DtoToModelLoanProgram(req))
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, models.Response{
		ID:     id,
		Status: "created",
	})
}

// @Router /admins/loan-programs/{code} [put]
// @Summary Обновление кредитной программы
// @Description Обновление кредитной программы
// @Tags programs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "токен авторизации"
// @Param code path string true "уникальный код программы"
// @Param program body dto.LoanProgramRequest true "кредитная программа"
// @Success 200
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
func (res AdminResource) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	code := chi.URLParam(r, "code")
	if code == "" {
		_ = render.Render(w, r, httperrors.BadRequest(fmt.Errorf("[ERROR] Empty Code")))
		return
	}

	var req dto.LoanProgramRequest

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err = req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	err = res.loanMan.UpdateLoanProgram(ctx, code, DtoToModelLoanProgram(req))
	switch err {
	case drivers.ErrDoesNotExist:
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case nil:
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

// @Router /admins/loan-programs/{code} [get]
// @Summary Получение кредитной программы по уникальному коду
// @Description Получение кредитной программы по уникальному коду
// @Tags programs
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "токен авторизации"
// @Param code path string true "уникальный код программы"
// @Success 200 {object} models.LoanProgram
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
func (res AdminResource) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := chi.URLParam(r, "code")

	loan, err := res.loanMan.LoanProgram(ctx, code)
	switch err {
	case drivers.ErrDoesNotExist:
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case nil:
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, loan)

}

// @Router /admins/loan-programs [get]
// @Summary Получение кредитных программ
// @Description Получение кредитных программ
// @Tags programs
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "токен авторизации"
// @Param page query int false "страница"
// @Param limit query int false "размер страницы"
// @Success 200 {object} dto.LoanPrograms
// @Failure 400 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
func (res AdminResource) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	page, limit, err := selector.ParsePaging(query)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	paging := selector.NewPaging().SetPaging(page, limit)

	programs, total, err := res.loanMan.LoanPrograms(ctx, selector.Paging{
		Skip:    paging.Skip,
		Limit:   paging.Limit,
		SortVal: -1,
		SortKey: "created_at",
	})
	switch err {
	case drivers.ErrDoesNotExist:
		_ = render.Render(w, r, httperrors.ResourceNotFound(err))
		return
	case nil:
	default:
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.LoanPrograms{
		Total:        total,
		LoanPrograms: programs,
	})
}
