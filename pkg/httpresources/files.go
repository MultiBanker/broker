package httpresources

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// FilesResource - для раздачи статичных файлов.
type FilesResource struct {
	path     string
	filesDir string
}

func NewFilesResource(path, filesDir string) *FilesResource {
	return &FilesResource{
		path:     path,
		filesDir: filesDir,
	}
}

func (res *FilesResource) Routes() chi.Router {
	r := chi.NewRouter()
	filesRoot := http.Dir(res.filesDir)

	newFileServer(r, "/", filesRoot)

	return r
}

func newFileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func (res *FilesResource) Path() string {
	return res.path
}

