package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/manager/user"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"
)

const tokenTTL = 60 * time.Minute

// @Summary Регистрация пользователя
// @Description Регистрирует пользователя в ССО
// @Accept json
// @Produce json
// @Tags auth
// @Param json body dto.SignUp true "Данные для регистрации пользователя"
// @Success 204
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/auth/signup [post]
func (res resource) signUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.SignUp

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	id, err := res.user.Create(ctx, models.User{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Patronymic: req.Patronymic,
		IIN:        req.IIN,
		Phone:      req.Phone,
		Password:   req.Password,
	})
	if err != nil {
		log.Printf("[ERROR] %v", err)
		_ = render.Render(w, r, httperrors.Internal(errors.New("server error")))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, models.Response{
		ID:     id,
		Status: "created",
	})

}

// @Summary Аутентификация по номеру телефона
// @Description Аутентифицирует пользователя по номеру телефона
// @Accept json
// @Produce json
// @Tags auth
// @Param json body dto.SignInByPhone true "Данные для быстрой аутентификации"
// @Success 200 {object} dto.NewJWTTokenResponse
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 410 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/auth/signin/phone [post]
func (res resource) signInByPhone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.SignInByPhone

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	userID, err := res.login.SignInByPhone(ctx, auth.NormPhoneNum(req.Phone), req.Password)
	if err != nil {
		if errors.Is(err, user.ErrUserPhoneIsNotVerified) {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		if errors.Is(err, user.ErrInvalidLoginOrPassword) {
			_ = render.Render(w, r, httperrors.BadRequest(err))
			return
		}
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	accessToken, err := res.auther.AccessToken(userID, "", "")
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	refreshToken, err := res.auther.RefreshToken(userID, "", "")
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.NewJWTTokenResponse{
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}

// @Summary Завершает сессию пользователя, удаляя JWT токен
// @Description Удаляет JWT токен в cookie.
// @Produce json
// @Tags auth
// @Security JWT
// @Param Authorization header string true "Токен аутентификации"
// @Success 200
// @Router /api/v1/users/auth/signout [delete]
func (res resource) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		HttpOnly: false,
		Expires:  time.Now().In(time.UTC).Add(-tokenTTL),
	})

	dto.RespondJSON(w, r, http.StatusOK, "JWT deleted successfully")
}
