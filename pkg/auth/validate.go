package auth

import "regexp"

const (
	PasswordLength = 8
	ru             = "ru"
	kz             = "kz"
	en             = "en"
)

var phoneRegex = regexp.MustCompile("^\\+?77([0124567][0-8]\\d{7})$")

func ValidatePhone(phone string) bool {
	return phoneRegex.MatchString(NormPhoneNum(phone))
}
