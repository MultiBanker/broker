package recovery

import (
	"encoding/json"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/servers/adminhttp/dto"
	_ "github.com/MultiBanker/broker/src/servers/clienthttp/dto"
	"github.com/go-chi/render"
)

// @Summary Отправка OTP на указанный номер телефона для восстановления пароля
// @Description Отправляет OTP на указанный номер телефона для восстановления пароля
// @Accept json
// @Produce json
// @Tags recovery
// @Param json body dto.RecoveryPhone true "Данные для восстановления пароля"
// @Success 204
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/recovery/phone [put]
func (res resource) recoveryPhone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.RecoveryPhone

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
// @Description Проверяет OTP при восстановлении пароля
// @Accept json
// @Produce json
// @Tags recovery
// @Param json body dto.RecoveryPhoneOTP true "Данные для валидации OTP"
// @Success 204
// @Failure 400 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/users/recovery/phone/otp [put]
func (res resource) recoveryPhoneOTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.RecoveryPhoneOTP

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	if err := req.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := res.user.ValidateOTP(ctx, "sms", req.Phone, req.OTP, req.Password); err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
