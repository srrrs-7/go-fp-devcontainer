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
)

func TestPutHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
		tests := []struct {
			name     string
			formData map[string]string
		}{
			{
				name: "valid request with all fields",
				formData: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Updated Task",
					"description": "Updated Description",
					"status":      "completed",
				},
			},
			{
				name: "valid request without description",
				formData: map[string]string{
					"id":     "550e8400-e29b-41d4-a716-446655440000",
					"title":  "Updated Task",
					"status": "pending",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					UpdateTaskFunc: func(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          arg.ID,
							Title:       arg.Title.String,
							Description: arg.Description,
							Status:      arg.Status.String,
						}, nil
					},
				}

				handler := NewPutHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if resp.StatusCode != http.StatusOK {
					t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
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
				name: "invalid uuid",
				formData: map[string]string{
					"id":    "invalid-uuid",
					"title": "Updated Task",
				},
			},
			{
				name: "missing id",
				formData: map[string]string{
					"title": "Updated Task",
				},
			},
			{
				name: "title too short",
				formData: map[string]string{
					"id":    "00000000-0000-0000-0000-000000000001",
					"title": "ab",
				},
			},
			{
				name: "missing title",
				formData: map[string]string{
					"id": "00000000-0000-0000-0000-000000000001",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					UpdateTaskFunc: func(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          arg.ID,
							Title:       arg.Title.String,
							Description: arg.Description,
							Status:      arg.Status.String,
						}, nil
					},
				}

				handler := NewPutHandler(mockDB)
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
