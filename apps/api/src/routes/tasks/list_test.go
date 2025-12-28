package tasks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils/db/db"
)

func TestListHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
		tests := []struct {
			name        string
			queryParams map[string]string
		}{
			{
				name: "valid request with all params",
				queryParams: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Test Task",
					"description": "Test Description",
					"status":      "pending",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
				q := req.URL.Query()
				for k, v := range tt.queryParams {
					q.Add(k, v)
				}
				req.URL.RawQuery = q.Encode()

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					ListTasksFunc: func(ctx context.Context) ([]db.Task, error) {
						return []db.Task{}, nil
					},
				}

				handler := NewListHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if resp.StatusCode != http.StatusOK {
					t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if _, ok := result["tasks"]; !ok {
					t.Error("expected 'tasks' field in response")
				}
			})
		}
	})

	t.Run("400 Bad Request", func(t *testing.T) {
		tests := []struct {
			name        string
			queryParams map[string]string
		}{
			{
				name: "invalid uuid",
				queryParams: map[string]string{
					"id":    "invalid",
					"title": "Test Task",
				},
			},
			{
				name: "title too short",
				queryParams: map[string]string{
					"id":    "00000000-0000-0000-0000-000000000001",
					"title": "ab",
				},
			},
			{
				name: "missing required id",
				queryParams: map[string]string{
					"title": "Test Task",
				},
			},
			{
				name: "missing required title",
				queryParams: map[string]string{
					"id": "00000000-0000-0000-0000-000000000001",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
				q := req.URL.Query()
				for k, v := range tt.queryParams {
					q.Add(k, v)
				}
				req.URL.RawQuery = q.Encode()

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					ListTasksFunc: func(ctx context.Context) ([]db.Task, error) {
						return []db.Task{}, nil
					},
				}

				handler := NewListHandler(mockDB)
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
