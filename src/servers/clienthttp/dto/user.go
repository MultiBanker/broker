package dto

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/hashicorp/go-multierror"
)

type UserApplication struct {
	UserID    string
	ChosenSKU string
}

func (ua UserApplication) Validate() error {
	var result *multierror.Error

	if ua.ChosenSKU == "" {
		result = multierror.Append(result, errors.New("chosen sku is empty"))
	}

	return result.ErrorOrNil()
}

type SignUp struct {
	FirstName  string `json:"first_name"` // Имя пользователя
	LastName   string `json:"last_name"`  // Фамилия пользователя
	Patronymic string `json:"patronymic"` // Отчество пользователя
	IIN        string `json:"iin"`
	Phone      string `json:"phone"`    // Номер телефона пользователя
	Password   string `json:"password"` // Пароль пользователя
}

func (su SignUp) Validate() error {
	var result *multierror.Error

	if su.FirstName == "" {
		result = multierror.Append(result, errors.New("first name is empty"))
	}
	if su.LastName == "" {
		result = multierror.Append(result, errors.New("last name is empty"))
	}
	if su.Patronymic == "" {
		result = multierror.Append(result, errors.New("patronymic is empty"))
	}
	if !auth.ValidatePhone(su.Phone) {
		result = multierror.Append(result, errors.New("phone is wrong"))
	}
	if utf8.RuneCountInString(su.Password) < auth.PasswordLength {
		result = multierror.Append(result, errors.New("password is wrong"))
	}

	if len(su.IIN) != 12 {
		result = multierror.Append(result, errors.New("iin. length must be 12"))
	}

	if !isDigit(su.IIN) {
		result = multierror.Append(result, errors.New("iin. must be digit"))
	}

	month, err := strconv.Atoi(su.IIN[2:4])
	if err != nil || month <= 0 || month > 12 {
		result = multierror.Append(result, errors.New("iin. invalid month in iin"))
	}

	day, err := strconv.ParseUint(su.IIN[4:6], 10, 64)
	if err != nil || day <= 0 || day > 31 {
		result = multierror.Append(result, errors.New("iin. invalid day in iin"))
	}

	return result.ErrorOrNil()
}

func isDigit(data string) bool {
	if _, err := strconv.Atoi(data); err != nil {
		return false
	}

	return true
}

// RecoveryPhone - модель для восстановления пароля зарегистрированного пользователя
type RecoveryPhone struct {
	Phone string `json:"phone"` // Номер телефона для восстановления
}

func (rp RecoveryPhone) Validate() error {
	var result *multierror.Error

	if !auth.ValidatePhone(rp.Phone) {
		result = multierror.Append(result, errors.New("invalid phone"))
	}

	return result.ErrorOrNil()
}

// VerifyPhone - модель для верификации номера телефона зарегистрированного пользователя
type VerifyPhone struct {
	Phone string `json:"phone"` // Номер телефона для верификации
}

func (vp VerifyPhone) Validate() error {
	var result *multierror.Error

	if !auth.ValidatePhone(vp.Phone) {
		result = multierror.Append(result, errors.New("invalid phone"))
	}

	return result.ErrorOrNil()
}

// RecoveryPhoneOTP - модель для подтверждения OTP отправленный по указанному номеру телефона
type RecoveryPhoneOTP struct {
	Phone    string `json:"phone"`    // Номер телефона для восстановления
	OTP      string `json:"otp"`      // One-Time-Password отправленный по СМС
	Password string `json:"password"` // Новый пароль пользователя
}

func (vp RecoveryPhoneOTP) Validate() error {
	var result *multierror.Error

	if !auth.ValidatePhone(vp.Phone) {
		result = multierror.Append(result, errors.New("invalid phone"))
	}

	if strings.TrimSpace(vp.OTP) == "" {
		result = multierror.Append(result, errors.New("empty otp"))
	}

	if utf8.RuneCountInString(vp.Password) < auth.PasswordLength {
		result = multierror.Append(result, errors.New("password is wrong"))
	}

	return result.ErrorOrNil()
}

// VerifyPhoneOTP - модель для подтверждения OTP отправленный по указанному номеру телефона
type VerifyPhoneOTP struct {
	Phone string `json:"phone"` // Номер телефона для верификации
	OTP   string `json:"otp"`   // One-Time-Password отправленный по СМС
}

// Validate валидация модели VerifyOTP
func (v VerifyPhoneOTP) Validate() error {
	var result *multierror.Error

	if !auth.ValidatePhone(v.Phone) {
		result = multierror.Append(result, errors.New("invalid phone"))
	}

	if strings.TrimSpace(v.OTP) == "" {
		result = multierror.Append(result, errors.New("empty otp"))
	}
	return result.ErrorOrNil()
}

// SignInByPhone - модель данных для аутентификации через номер телефона и пароль
type SignInByPhone struct {
	Phone    string `json:"phone"`    // Номер телефона для аутентификации
	Password string `json:"password"` // Пароль для аутентификации
}

func (vp SignInByPhone) Validate() error {
	var result *multierror.Error

	if !auth.ValidatePhone(vp.Phone) {
		result = multierror.Append(result, errors.New("invalid phone"))
	}

	if utf8.RuneCountInString(vp.Password) < auth.PasswordLength {
		result = multierror.Append(result, errors.New("password is wrong"))
	}

	return result.ErrorOrNil()
}

type NewJWTTokenResponse struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
