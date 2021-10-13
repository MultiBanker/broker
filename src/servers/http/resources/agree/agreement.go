package agree

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/http/detector"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// create godoc
// @Tags Agreements
// @Summary Создание новой спецификации для соглашения
// @Description Создание новой спецификации для соглашения
// @ID specification-create
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param specification body models.Specification true "specification"
// @Success 201 {object} dto.IDResponse
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 412 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements [post]
func (res Resource) createAgree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var specification models.Specification
	if err := json.NewDecoder(r.Body).Decode(&specification); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.specMan.Create(ctx, specification)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, dto.IDResponse{
		ID:     id,
		Status: "created",
	})
}

// get godoc
// @Summary Получение соглашений
// @Description Получение соглашений
// @Tags Agreements
// @ID specification-get
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the specification"
// @Success 200 {object} models.Specification
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 412 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/{id} [get]
func (res Resource) getAgree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	specification, err := res.specMan.Get(ctx, id)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.JSON(w, r, specification)
}

// get godoc
// @Summary Получить соглашение по коду
// @Description Получить соглашение по коду
// @Tags Agreements
// @ID specification-get-by-code
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param code path string true "code of the specification"
// @Success 200 {object} models.Specification
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 412 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/{code} [get]
func (res Resource) getByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := chi.URLParam(r, "code")
	specification, err := res.specMan.GetByCode(ctx, code)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.JSON(w, r, specification)
}

// list godoc
// @Summary Вернуть все соглашения
// @Description Вернуть все соглашения
// @Tags Agreements
// @ID specification-root
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success 200 {array} models.Specification
// @Failure 400 {string} httperrors.Response
// @Failure 404 {string} httperrors.Response
// @Failure 500 {string} httperrors.Response
// @Router /agreements [get]
func (res Resource) getAgrees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paging := selector.NewPaging()
	paging.SetSorting("created_at", -1)

	q := r.URL.Query()

	skipStr := q.Get("skip")
	if len(skipStr) == 0 {
		paging.Skip = defaultSkip
	} else {
		skip, err := strconv.ParseInt(skipStr, 10, 64)
		if err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		if skip < 0 {
			_ = render.Render(w, r, httperrors.BadRequest(errors.New("skip cant be less than 0")))
			return
		}
		paging.Skip = skip
	}

	limitStr := q.Get("limit")
	if len(limitStr) == 0 {
		paging.Limit = defaultLimit
	} else {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		if limit <= 0 && limit > maxLimit {
			_ = render.Render(w, r, httperrors.BadRequest(errors.New(fmt.Sprintf("limit must be between 1 and %v", maxLimit))))
			return
		}
		paging.Limit = limit
	}

	specifications, count, err := res.specMan.List(ctx, *paging)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.Specifications{
		Total:          count,
		Specifications: specifications,
	})
}

// update godoc
// @Summary Обновление соглашений
// @Description Обновление соглашений
// @Tags Agreements
// @ID specification-update
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the specification"
// @Param specifications body models.Specification true "specification"
// @Success 204
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 412 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/{id} [PUT]
func (res Resource) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	var specification models.Specification
	if err := json.NewDecoder(r.Body).Decode(&specification); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	specification.ID = id
	err := res.specMan.Update(ctx, specification)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.NoContent(w, r)
}

