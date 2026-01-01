package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils/db/db"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestGetHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
		type expected struct {
			statusCode int
			hasID      bool
		}

		tests := []struct {
			name        string
			queryParams map[string]string
			expected    expected
		}{
			{
				name: "valid uuid",
				queryParams: map[string]string{
					"id": "550e8400-e29b-41d4-a716-446655440000",
				},
				expected: expected{
					statusCode: http.StatusOK,
					hasID:      true,
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
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				_, hasID := result["id"]
				if diff := cmp.Diff(tt.expected.hasID, hasID); diff != "" {
					t.Errorf("'id' field presence mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("400 Bad Request", func(t *testing.T) {
		type expected struct {
			statusCode int
			hasType    bool
		}

		tests := []struct {
			name        string
			queryParams map[string]string
			expected    expected
		}{
			{
				name: "invalid uuid",
				queryParams: map[string]string{
					"id": "invalid-uuid",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name:        "missing id",
				queryParams: map[string]string{},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "empty id",
				queryParams: map[string]string{
					"id": "",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
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
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				_, hasType := result["type"]
				if diff := cmp.Diff(tt.expected.hasType, hasType); diff != "" {
					t.Errorf("'type' field presence mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}
