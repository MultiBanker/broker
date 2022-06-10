package verification

import (
	"encoding/json"
	"net/http"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	_ "github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"
)

// @Summary Отправка OTP на указанный номер телефона для верификации
// @Description Отправляет OTP на указанный номер телефона для верификации
// @Accept json
// @Produce json
// @Tags verify
// @Param json body dto.VerifyPhone true "Данные для верификации телефона"
// @Success 204
// @Failure 400 {object} httperrors.Response
// @Failure 410 {object} httperrors.Response
// @Failure 429 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/verify/phone [put]
func (res resource) verifyPhone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.VerifyPhone

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := res.user.SendOTP(ctx, "sms", req.Phone); err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

// @Summary Валидация OTP по номеру телефона
// @Description Проверяет OTP при верификации номера телефона и выдает JWT
// @Accept json
// @Produce json
// @Tags verify
// @Param json body dto.VerifyPhoneOTP true "Данные для валидации OTP"
// @Success 200 {object} dto.NewJWTTokenResponse
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/verify/phone/otp [put]
func (res resource) verifyPhoneOTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.VerifyPhoneOTP

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	userID, err := res.user.ValidateOTP(ctx, "sms", auth.NormPhoneNum(req.Phone), req.OTP)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
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
