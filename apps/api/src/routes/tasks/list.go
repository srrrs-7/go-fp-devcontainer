package tasks

import (
	"net/http"
)

type listResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ListHandler(w http.ResponseWriter, r *http.Request) {}
