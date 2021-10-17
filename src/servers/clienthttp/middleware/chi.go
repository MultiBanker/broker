package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const file = "/files"

func Mount(version, filesDir, basePath string) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.NoCache) // no-cache
	//r.Use(middleware.RequestID) // вставляет request ID в контекст каждого запроса
	r.Use(middleware.Logger)    // логирует начало и окончание каждого запроса с указанием времени обработки
	r.Use(middleware.Recoverer) // управляемо обрабатывает паники и выдает stack trace при их возникновении
	r.Use(middleware.RealIP)    // устанавливает RemoteAddr для каждого запроса с заголовками X-Forwarded-For или X-Real-IP
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	// монтируем дополнительные ресурсы
	r.Mount("/version", VersionResource{Version: version}.Routes())
	r.Mount(file, FilesResource{FilesDir: filesDir}.Routes())
	r.Mount("/swagger", SwaggerResource{BasePath: basePath, FilesPath: file}.Routes())

	return r
}
