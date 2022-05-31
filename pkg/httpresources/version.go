package httpresources

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const APIVersion = "v1"

type VersionResponse struct {
	Swagger string `json:"swagger"`
	Version string `json:"version"`
}

type VersionResource struct {
	path    string
	Version string
}

func NewVersionResource(path, version string) *VersionResource {
	return &VersionResource{
		path:    path,
		Version: version,
	}
}

func (res *VersionResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", res.Get)

	return r
}

func (res *VersionResource) Get(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, VersionResponse{
		Swagger: APIVersion,
		Version: res.Version,
	})
}

func (res *VersionResource) Path() string {
	return res.path
}

