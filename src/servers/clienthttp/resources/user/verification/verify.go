package verification

import (
	"net/http"

	_ "github.com/MultiBanker/broker/src/servers/clienthttp/dto"
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
func (res resource) verifyPhone(w http.ResponseWriter, r *http.Request) {}


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
func (res resource) verifyPhoneOTP(w http.ResponseWriter, r *http.Request) {}