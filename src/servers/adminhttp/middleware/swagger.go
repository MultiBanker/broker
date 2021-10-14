package middleware

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

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

	r.Group(func(r chi.Router) {
		r.Get("/admin/*", sr.Indexer("admin"))
	})

	return r
}

func (sr SwaggerResource) Indexer(role string) http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL(filepath.Join(sr.BasePath, sr.FilesPath, role, "swagger.json")),
	)
}