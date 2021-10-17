package guard

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

var (
	ErrNoRoles     = errors.New("no roles found")
	ErrForbidden   = errors.New("forbidden")
	ErrInvalidType = errors.New("invalid type")
)

// Allow middleware guards endpoints from certain user roles
func Allow(errs ErrorRenderer, allow ...roleName) func(handler http.Handler) http.Handler {
	if errs == nil {
		errs = DefaultRender
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := Allowed(r.Context(), allow...); err != nil {
				errs(w, r, err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Allowed checks if context contains allowed role
func Allowed(ctx context.Context, allow ...roleName) error {
	token, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return err
	}

	err = jwt.Validate(token)
	if err != nil {
		return err
	}

	tempRoles, ok := claims["roles"].([]interface{})
	if !ok {
		return ErrNoRoles
	}

	roles := make([]roleName, 0, len(tempRoles))
	for i := range tempRoles {
		role, ok := tempRoles[i].(string)
		if !ok {
			return fmt.Errorf("%w: expected string", ErrInvalidType)
		}
		roles = append(roles, roleName(role))
	}

	var allowed bool
	for i := range allow {
		allowed = allowed || check(allow[i], roles...)
	}

	if !allowed {
		return ErrForbidden
	}

	return nil
}
