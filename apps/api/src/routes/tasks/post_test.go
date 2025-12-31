package tasks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"utils/db/db"

	"github.com/google/uuid"
)

func TestPostHandler(t *testing.T) {
	t.Run("201 Created", func(t *testing.T) {
		tests := []struct {
			name     string
			formData map[string]string
		}{
			{
				name: "valid request with title and description",
				formData: map[string]string{
					"title":       "New Task",
					"description": "Task Description",
				},
			},
			{
				name: "valid request with title only",
				formData: map[string]string{
					"title": "New Task",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					CreateTaskFunc: func(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          uuid.New(),
							Title:       arg.Title,
							Description: arg.Description,
							Status:      "pending",
						}, nil
					},
				}

				handler := NewPostHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if resp.StatusCode != http.StatusCreated {
					t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if _, ok := result["id"]; !ok {
					t.Error("expected 'id' field in response")
				}
			})
		}
	})

	t.Run("400 Bad Request", func(t *testing.T) {
		tests := []struct {
			name     string
			formData map[string]string
		}{
			{
				name: "title too short",
				formData: map[string]string{
					"title": "ab",
				},
			},
			{
				name: "missing title",
				formData: map[string]string{
					"description": "Only description",
				},
			},
			{
				name: "empty title",
				formData: map[string]string{
					"title": "",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					CreateTaskFunc: func(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          uuid.New(),
							Title:       arg.Title,
							Description: arg.Description,
							Status:      "pending",
						}, nil
					},
				}

				handler := NewPostHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if resp.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if _, ok := result["type"]; !ok {
					t.Error("expected 'type' field in error response")
				}
			})
		}
	})
}
