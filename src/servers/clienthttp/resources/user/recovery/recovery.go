package recovery

import (
	"net/http"

	_ "github.com/MultiBanker/broker/src/servers/clienthttp/dto"
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
func (res resource) recoveryPhone(w http.ResponseWriter, r *http.Request) {}

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
func (res resource) recoveryPhoneOTP(w http.ResponseWriter, r *http.Request) {}
