package tasks

import "net/http"

type postResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {}
