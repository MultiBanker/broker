package httpresources

import (
	"path/filepath"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//SwaggerResource - для размещения API документации
type SwaggerResource struct {
	mountPath string
	basePath  string
	filesPath string
}

func NewSwaggerResource(mountPath, basePath string, filesPath string) *SwaggerResource {
	return &SwaggerResource{
		mountPath: mountPath,
		basePath:  basePath,
		filesPath: filesPath,
	}
}

func (res *SwaggerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL(filepath.Join(res.basePath, res.filesPath, "swagger.json")),
	))

	return r
}

func (res *SwaggerResource) Path() string {
	return res.mountPath
}

