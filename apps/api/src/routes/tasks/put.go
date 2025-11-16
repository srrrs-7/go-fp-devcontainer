package tasks

import "net/http"

type putResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {}
