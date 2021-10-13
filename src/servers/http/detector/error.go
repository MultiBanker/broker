package detector

import (
	"errors"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/go-chi/render"
)

func Error(err error) render.Renderer {
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		return httperrors.ResourceNotFound(err)
	}

	return httperrors.Internal(err)
}

