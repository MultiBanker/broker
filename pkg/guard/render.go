package guard

import "net/http"

type ErrorRenderer func(w http.ResponseWriter, r *http.Request, err error)

func DefaultRender(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte(err.Error()))
}

