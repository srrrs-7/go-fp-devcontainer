package tasks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils/db/db"

	"github.com/google/go-cmp/cmp"
)

func TestListHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
		type expected struct {
			statusCode int
			hasTasks   bool
		}

		tests := []struct {
			name        string
			queryParams map[string]string
			expected    expected
		}{
			{
				name: "valid request with all params",
				queryParams: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Test Task",
					"description": "Test Description",
					"status":      "pending",
				},
				expected: expected{
					statusCode: http.StatusOK,
					hasTasks:   true,
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
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				_, hasTasks := result["tasks"]
				if diff := cmp.Diff(tt.expected.hasTasks, hasTasks); diff != "" {
					t.Errorf("'tasks' field presence mismatch (-want +got):\n%s", diff)
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
					"id":    "invalid",
					"title": "Test Task",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "title too short",
				queryParams: map[string]string{
					"id":    "00000000-0000-0000-0000-000000000001",
					"title": "ab",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "missing required id",
				queryParams: map[string]string{
					"title": "Test Task",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "missing required title",
				queryParams: map[string]string{
					"id": "00000000-0000-0000-0000-000000000001",
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
					ListTasksFunc: func(ctx context.Context) ([]db.Task, error) {
						return []db.Task{}, nil
					},
				}

				handler := NewListHandler(mockDB)
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
