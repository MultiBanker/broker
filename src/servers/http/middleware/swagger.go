package middleware

import (
	"path/filepath"

	"github.com/go-chi/chi"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerResource для размещения API документации
type SwaggerResource struct {
	BasePath  string
	FilesPath string
}

func (sr SwaggerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL(filepath.Join(sr.BasePath, sr.FilesPath, "swagger.json")),
	))
	return r
}
