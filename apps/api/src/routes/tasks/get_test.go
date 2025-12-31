package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils/db/db"

	"github.com/google/uuid"
)

func TestGetHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
		tests := []struct {
			name        string
			queryParams map[string]string
		}{
			{
				name: "valid uuid",
				queryParams: map[string]string{
					"id": "550e8400-e29b-41d4-a716-446655440000",
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
					GetTaskFunc: func(ctx context.Context, id uuid.UUID) (db.Task, error) {
						return db.Task{
							ID:          id,
							Title:       "Sample Task",
							Description: sql.NullString{String: "Description", Valid: true},
							Status:      "pending",
						}, nil
					},
				}

				handler := NewGetHandler(mockDB)
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
			name        string
			queryParams map[string]string
		}{
			{
				name: "invalid uuid",
				queryParams: map[string]string{
					"id": "invalid-uuid",
				},
			},
			{
				name:        "missing id",
				queryParams: map[string]string{},
			},
			{
				name: "empty id",
				queryParams: map[string]string{
					"id": "",
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
					GetTaskFunc: func(ctx context.Context, id uuid.UUID) (db.Task, error) {
						return db.Task{
							ID:          id,
							Title:       "Sample Task",
							Description: sql.NullString{String: "Description", Valid: true},
							Status:      "pending",
						}, nil
					},
				}

				handler := NewGetHandler(mockDB)
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
