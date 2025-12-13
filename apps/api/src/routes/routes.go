package routes

import (
	"api/src/routes/tasks"
	"database/sql"
	"net/http"
	"utils/db/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(conn *sql.DB) http.Handler {
	r := chi.NewRouter()
	q := db.New(conn)

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
				r.Get("/", tasks.NewListHandler(q))
				r.Post("/", tasks.NewPostHandler(q))
				r.Get("/{id}", tasks.NewGetHandler(q))
				r.Put("/{id}", tasks.NewPutHandler(q))
			})
		})
	})

	return r
}
