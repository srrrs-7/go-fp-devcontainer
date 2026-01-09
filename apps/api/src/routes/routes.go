package routes

import (
	mw "api/src/routes/middleware"
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
	r.Use(middleware.Recoverer)
	r.Use(mw.Logger)

	// health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// Bearer認証を適用
			r.Use(mw.BearerAuth(validateToken))

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

// validateToken はトークンを検証する（TODO: 実際の検証ロジックに置き換え）
func validateToken(token string) (bool, error) {
	// TODO: JWT検証やCognito検証などを実装
	return token != "", nil
}
