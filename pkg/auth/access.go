package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
)

var ErrInvalidToken = errors.New("token is incorrect or expired")

type ctxKey int

const (
	UserIDKey ctxKey = iota + 1
	RolesKey
)

type UserAccessCtx struct {
	jwtauth *jwtauth.JWTAuth
}

func NewUserAccessCtx(jwtKey []byte) *UserAccessCtx {
	return &UserAccessCtx{
		jwtauth: jwtauth.New("HS256", jwtKey, nil),
	}
}

func (a *UserAccessCtx) Middleware(next http.Handler) http.Handler {
	return jwtauth.Verifier(a.jwtauth)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		err = jwt.Validate(token)
		if err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		if claims["is_refresh"].(bool) {
			_ = render.Render(w, r, httperrors.Unauthorized(ErrInvalidToken))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims["user_id"])
		ctx = context.WithValue(ctx, RolesKey, claims["roles"])

		// токен валидный, пропускаем его
		next.ServeHTTP(w, r.WithContext(ctx))
	}))
}
