package middleware

import (
	"context"
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/MultiBanker/broker/src/models"
)

type ctxKey int

const (
	IDKey ctxKey = iota + 1
	RolesKey
)

type UserAccessCtx struct {
	jwtKey []byte
}

func NewUserAccessCtx(jwtKey []byte) *UserAccessCtx {
	return &UserAccessCtx{
		jwtKey: jwtKey,
	}
}

func (ua *UserAccessCtx) ChiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		if err = jwt.Validate(token); err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		if isRefresh, ok := claims["is_refresh"].(bool); !ok || isRefresh {
			_ = render.Render(w, r, httperrors.Unauthorized(models.ErrInvalidToken))

			return
		}

		ctx := r.Context()

		ID, hasID := claims["id"]
		if !hasID {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		ctx = context.WithValue(ctx, IDKey, ID)
		ctx = context.WithValue(ctx, RolesKey, claims["roles"])

		// токен валидный, пропускаем его
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PermissionMiddleware - берет роли из контекста и проверяет доступ к ресурсу по ролям юзера. Проверяет наличие роли юзера в слайсе allowedRoles
func (ua *UserAccessCtx) PermissionMiddleware(allowedRoles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roles, ok := r.Context().Value("roles").([]string)
			if !ok {
				_ = render.Render(w, r, httperrors.AccessDenied(models.ErrInvalidRole))
				return
			}

			allowedRolesMap := make(map[string]struct{}, len(allowedRoles))
			for _, role := range allowedRoles {
				allowedRolesMap[role] = struct{}{}
			}

			for _, role := range roles {
				if _, ok = allowedRolesMap[role]; ok {
					next.ServeHTTP(w, r)
					return
				}
			}

			_ = render.Render(w, r, httperrors.AccessDenied(models.ErrInvalidRole))
		})
	}
}
