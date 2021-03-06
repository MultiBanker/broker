package auth

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type Authenticator interface {
	AccessToken(id, role, code string) (string, error)
	RefreshToken(id, role, code string) (string, error)
	// code, type
	Tokens(id, role, code string) (access string, refresh string, err error)
	TokenAuth() *jwtauth.JWTAuth
	JWTKey() []byte
}

var _ Authenticator = (*Authenticate)(nil)

type Authenticate struct {
	jwtKey          []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewAuthenticator(jwtKey []byte, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *Authenticate {
	return &Authenticate{jwtKey: jwtKey, AccessTokenTTL: accessTokenTTL, RefreshTokenTTL: refreshTokenTTL}
}

// JWT key
func (a *Authenticate) JWTKey() []byte {
	return a.jwtKey
}

func (a *Authenticate) TokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", a.jwtKey, nil)
}

func (a *Authenticate) AccessToken(id, role, code string) (string, error) {
	claims := claimer(id, role, code, a.AccessTokenTTL)
	claims["is_refresh"] = false
	_, token, err := a.TokenAuth().Encode(claims)
	return token, err
}

func (a *Authenticate) RefreshToken(id, role, code string) (string, error) {
	claims := claimer(id, role, code, a.RefreshTokenTTL)
	claims["is_refresh"] = true
	_, token, err := a.TokenAuth().Encode(claims)
	return token, err
}

func (a Authenticate) Tokens(id, role, code string) (access string, refresh string, err error) {
	access, err = a.AccessToken(id, role, code)
	if err != nil {
		return "", "", err
	}
	refresh, err = a.RefreshToken(id, role, code)
	return access, refresh, err
}

func claimer(id, role, code string, duration time.Duration) map[string]interface{} {
	return map[string]interface {
	}{
		"user_id":    id,
		"role":       role,
		"code":       code,
		"expired_at": time.Now().Add(duration).Unix(),
	}
}

