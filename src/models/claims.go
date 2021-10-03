package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims структура, хранящая закодированный JWT авторизации.
// Встраивание для предоставления поля expiry.
type Claims struct {
	ID        string   `json:"id"`
	IsRefresh bool     `json:"is_refresh"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}
