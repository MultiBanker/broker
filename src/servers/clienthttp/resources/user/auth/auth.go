package auth

import (
	"net/http"
	"time"

	"github.com/MultiBanker/broker/src/servers/clienthttp/dto"
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
func (res resource) signUp(w http.ResponseWriter, r *http.Request) {}

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
func (res resource) signInByPhone(w http.ResponseWriter, r *http.Request) {}

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
