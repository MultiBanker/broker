package auth

import "github.com/dongri/phonenumber"

// NormPhoneNum нормализует телефонные номера.
func NormPhoneNum(phone string) string {
	if len(phone) > 10 && phone[0] == '8' {
		return phonenumber.Parse(phone[1:], "KZ")
	}

	return phonenumber.Parse(phone, "KZ")
}

// MaskPhoneNum накладывает скрывающую маску на номер телефона
func MaskPhoneNum(phone string) string {

	phone = NormPhoneNum(phone)
	maskedPhone := ""
	for i, c := range phone {
		if i >= 4 && i <= 8 {
			maskedPhone += "*"
			continue
		}
		maskedPhone += string(c)
	}

	return maskedPhone
}

