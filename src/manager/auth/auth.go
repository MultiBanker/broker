package auth

import (
	"time"

	"github.com/MultiBanker/broker/src/config"
	"github.com/go-chi/jwtauth"
)

type Authenticator interface {
	AccessToken(id string) (string, error)
	RefreshToken(id string) (string, error)
	Tokens(id string) (access string, refresh string, err error)
	TokenAuth() *jwtauth.JWTAuth
	JWTKey() []byte
}

type Authenticate struct {
	jwtKey          []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewAuthenticator(opts *config.Config) Authenticator {
	return &Authenticate{
		jwtKey:          []byte(opts.JWTKey),
		AccessTokenTTL:  time.Duration(opts.AccessToken)*time.Hour,
		RefreshTokenTTL: time.Duration(opts.RefreshToken)*24*time.Hour*30,
	}
}

// JWT key
func (a *Authenticate) JWTKey() []byte {
	return a.jwtKey
}

func (a *Authenticate) TokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", a.jwtKey, nil)
}

func (a *Authenticate) AccessToken(id string) (string, error) {
	claims := claimer(id, a.AccessTokenTTL)
	_, token, err := a.TokenAuth().Encode(claims)
	return token, err
}

func (a *Authenticate) RefreshToken(id string) (string, error) {
	claims := claimer(id, a.RefreshTokenTTL)
	_, token, err := a.TokenAuth().Encode(claims)
	return token, err
}

func (a Authenticate) Tokens(id string) (access string, refresh string, err error) {
	access, err = a.AccessToken(id)
	if err != nil {
		return "", "", err
	}
	refresh, err = a.RefreshToken(id)
	return access, refresh, err
}

func claimer(id string, duration time.Duration) map[string]interface{} {
	return map[string]interface {
	}{
		"id":         id,
		"expired_at": time.Now().Add(duration).Unix(),
	}
}
