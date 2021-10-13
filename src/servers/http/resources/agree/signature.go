package agree

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"github.com/MultiBanker/broker/src/servers/http/detector"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	defaultSkip  int64 = 0
	defaultLimit int64 = 10
	maxLimit     int64 = 100
)

// @Tags Signature
// @Summary Создание новой подписи
// @Description Создание новой подписи
// @Accept  json
// @Produce  json
// @Param signature body dto.CreateSignatureReq true "body"
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.IDResponse
// @Failure 400 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/signatures [post]
func (res Resource) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateSignatureReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.signatureMan.Create(ctx, req)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.IDResponse{
		ID:     id,
		Status: "created",
	})

}

// get godoc
// @Summary Получение подписи
// @Description Получение подписи
// @Tags Signature
// @ID signature-get
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path string true "id of the signature"
// @Success 200 {object} models.Signature
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 412 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/signatures/{id} [get]
func (res *Resource) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	s, err := res.signatureMan.Get(ctx, id)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.JSON(w, r, s)
}

// list godoc
// @Summary Получение подписей
// @Description Получение подписей
// @Tags Signature
// @ID signature-list
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int false "pagination limit"
// @Param skip query int false "pagination skip"
// @Success 200 {object} dto.Signature
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /agreements/signatures [get]
func (res *Resource) list(w http.ResponseWriter, r *http.Request) {
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

	signatures, count, err := res.signatureMan.List(ctx, *paging)
	if err != nil {
		log.Println("[ERROR]", err)
		_ = render.Render(w, r, detector.Error(err))
		return
	}

	render.JSON(w, r, dto.Signature{
		Total:     count,
		Signature: signatures,
	})
}
