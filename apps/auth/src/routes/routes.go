package routes

import (
	"api/src/routes/tasks"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// Tasks
			r.Route("/tasks", func(r chi.Router) {
				r.Get("/", tasks.ListHandler)
				r.Post("/", tasks.PostHandler)
				r.Get("/{id}", tasks.GetHandler)
				r.Put("/{id}", tasks.PutHandler)
			})
		})
	})

	return r
}
